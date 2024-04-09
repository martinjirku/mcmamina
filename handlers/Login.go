package handlers

import (
	"log/slog"
	"net/http"

	"jirku.sk/mcmamina/template/components"
	"jirku.sk/mcmamina/template/pages"
)

func Login(Log *slog.Logger, cssPathGetter CSSPathGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		cssPath, _ := cssPathGetter.GetCssPath()
		dto := pages.NewLoginPageDto("", "", "")
		// TODO: handle GET
		// TODO: get CSRF token
		components.Page(components.NewPage(
			"Login",
			"Prihlásenie",
			cssPath,
			pages.LoginPage(dto),
		)).Render(r.Context(), w)
	}
}
