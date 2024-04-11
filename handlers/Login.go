package handlers

import (
	"log/slog"
	"net/http"

	"github.com/justinas/nosurf"
	"jirku.sk/mcmamina/template/components"
	"jirku.sk/mcmamina/template/pages"
)

func Login(Log *slog.Logger,
	cssPathGetter CSSPathGetter,
	recaptcha RecaptchaValidator,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		page := LoginPage{
			cssPathGetter: cssPathGetter,
			recaptcha:     recaptcha,
			log:           Log,
		}

		switch r.Method {
		case http.MethodGet:
			page.loginGet(w, r)
		case http.MethodPost:
			page.loginAction(w, r)
		}
	}
}

type LoginPage struct {
	cssPathGetter CSSPathGetter
	recaptcha     RecaptchaValidator
	log           *slog.Logger
}

func (l *LoginPage) loginGet(w http.ResponseWriter, r *http.Request) {
	cssPath, _ := l.cssPathGetter.GetCssPath()
	dto := pages.NewLoginPageDto(nosurf.Token(r), "", l.recaptcha.Key())
	components.Page(components.NewPage(
		"Login",
		"Prihlásenie",
		cssPath,
		pages.LoginPage(dto),
	)).Render(r.Context(), w)
}
func (l *LoginPage) loginAction(w http.ResponseWriter, r *http.Request) {
	cssPath, _ := l.cssPathGetter.GetCssPath()
	r.ParseForm()
	username := r.Form.Get("username")

	err := l.recaptcha.ValidateCaptcha(r)
	dto := pages.NewLoginPageDto(nosurf.Token(r), username, l.recaptcha.Key())
	if err != nil {
		l.log.Info("recaptcha validation failed", slog.Any("error", err))
		dto.SetErrorMsg("Chyba pri overovaní reCAPTCHA")
	} else {
		dto.SetErrorMsg("Nesprávne prihlasovacie údaje")
	}
	components.Page(components.NewPage(
		"Login",
		"Prihlásenie",
		cssPath,
		pages.LoginPage(dto),
	)).Render(r.Context(), w)
}
