package main

import (
		"context"
		"log"
		"net/http"
		"time"
		"os"

		"qrsink/views"

		"github.com/joho/godotenv"
		"github.com/go-chi/chi/v5"
		"github.com/redis/go-redis/v9"
)

type ScanResult struct {
		Code      string
		Status    string
		Timestamp time.Time
}

var rdb *redis.Client

func main() {

		if err := godotenv.Load(); err != nil {
				log.Fatal("Error loading .env file")
		}

		rdb = redis.NewClient(&redis.Options{
				Addr:     os.Getenv("REDIS_ADDR"),
				Username: os.Getenv("REDIS_USER"),
				Password: os.Getenv("REDIS_PASS"),
				DB:       0,
		})

		if err := rdb.Ping(context.Background()).Err(); err != nil {
				log.Fatal("Redis connection failed:", err)
		}

		r := chi.NewRouter()

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "/entry", http.StatusSeeOther)
		})
		r.Get("/entry", showEntryForm)
		r.Get("/scan/{code}", handleScan)
		r.Post("/result", validateManualEntry)

		port := os.Getenv("PORT")
		if port == "" {
				port = "8080"
		}
		log.Fatal(http.ListenAndServe(":" + port, r))
}

func showEntryForm(w http.ResponseWriter, r *http.Request) {
		views.Index().Render(r.Context(), w)
}

func handleScan(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		result := processCode(code)
		renderResult(w, result)
}

func validateManualEntry(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		result := processCode(code)
		renderResult(w, result)
}

func processCode(code string) ScanResult {
		ctx := context.Background()

		exists, err := rdb.Exists(ctx, code).Result()
		if err != nil {
				return ScanResult{Code: code, Status: "error", Timestamp: time.Now()}
		}

		if exists == 1 {
				return ScanResult{Code: code, Status: "duplicate", Timestamp: time.Now()}
		}

		if !isValidCode(code) {
				return ScanResult{Code: code, Status: "invalid", Timestamp: time.Now()}
		}

		err = rdb.Set(ctx, code, time.Now().String(), 0).Err()
		if err != nil {
				return ScanResult{Code: code, Status: "error", Timestamp: time.Now()}
		}

		return ScanResult{Code: code, Status: "valid", Timestamp: time.Now()}
}

func renderResult(w http.ResponseWriter, result ScanResult) {
		views.Result(views.ResultData{
				Code:      result.Code,
				Status:    result.Status,
				Timestamp: result.Timestamp,
		}).Render(context.Background(), w)
}

func isValidCode(code string) bool {
		if len(code) != 8 {
				return false
		}
		return true
}
