// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func layout(title string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html><head><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/layout.templ`, Line: 7, Col: 25}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><style>\n                :root {\n                    --primary: #FF3366;\n                    --bg: #111111;\n                    --text: #FFFFFF;\n                }\n\n                body {\n                    background: var(--bg);\n                    color: var(--text);\n                    font-family: system-ui;\n                    min-height: 100vh;\n                    margin: 0;\n                    display: grid;\n                    place-items: center;\n                }\n\n                .container {\n                    text-align: center;\n                    padding: 2rem;\n                    max-width: 500px;\n                    width: 100%;\n                }\n\n                input {\n                    width: 100%;\n                    padding: 1rem;\n                    font-size: 1.25rem;\n                    border: 2px solid var(--primary);\n                    border-radius: 8px;\n                    background: transparent;\n                    color: var(--text);\n                    margin-bottom: 1rem;\n                    transition: all 0.3s ease;\n                }\n\n                input:focus {\n                    outline: none;\n                    box-shadow: 0 0 15px var(--primary);\n                }\n\n                button {\n                    background: var(--primary);\n                    color: var(--text);\n                    border: none;\n                    padding: 1rem 2rem;\n                    font-size: 1.25rem;\n                    border-radius: 8px;\n                    cursor: pointer;\n                    transition: transform 0.2s ease;\n                }\n\n                button:hover {\n                    transform: translateY(-2px);\n                }\n\n                .result {\n                    margin-top: 2rem;\n                }\n\n                .valid {\n                    color: #4CAF50;\n                }\n\n                .invalid, .error {\n                    color: #f44336;\n                }\n\n                .duplicate {\n                    color: #FFC107;\n                }\n            </style></head><body><div class=\"container\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
