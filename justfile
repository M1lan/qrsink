#!/usr/bin/env -S just --justfile

doit:
  go mod tidy
  templ generate
  go build .
  echo "Build successful"
  echo "You can run ./qrsink now and then visit http://127.0.0.1:8080"
