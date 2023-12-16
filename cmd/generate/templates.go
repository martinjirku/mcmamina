package main

const pageTemplTmpl = `
package pages

import (
	"jirku.sk/mcmamina/template/layout"
	"jirku.sk/mcmamina/template/components"
)

templ {{.Name}}Page() {
	@layout.Layout(templ.CSSClasses{"{{.ClassName }} w-full bg-cover bg-center text-indigo-800 font-light"}, func(link string) bool { return link == "{{.Path}}"}) {
		@components.FullWidthCard(components.NewFullWidthCard().Margin("mb-0")) {
			@components.CardContent("") {
				<div>{{.Title}} Page</div>
			}
		}
	}
}
`
const pageTs = `
import "./{{.PageCSSName}}"
`
const pageCss = `
.{{.ClassName }} {
	background-image: url("@assets/images/crayons-1445053_640.jpg");
}
`

const pageHandlerGo = `
package handlers

import (
	"log/slog"
	"net/http"

	"jirku.sk/mcmamina/template/components"
	"jirku.sk/mcmamina/template/pages"
)

func {{.Name}}(Log *slog.Logger, cssPathGetter CSSPathGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		cssPath, _ := cssPathGetter.GetCssPath()
		// TODO: handle GET
		components.Page(components.NewPage(
			"{{.Name}}",
			"{{.Title}}",
			cssPath,
			pages.{{.Name}}Page(),
		)).Render(r.Context(), w)
	}
}
`
