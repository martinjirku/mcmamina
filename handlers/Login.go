package handlers

import (
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/justinas/nosurf"
)

type RecaptchaValidator interface {
	ValidateCaptcha(r *http.Request) error
	Key() string
}

type LoginPage struct {
	cssPathGetter CSSPathGetter
	recaptcha     RecaptchaValidator
	log           *slog.Logger
	tmpl          *template.Template
}

func Login(log *slog.Logger, cssPathGetter CSSPathGetter, recaptcha RecaptchaValidator, tmpl *template.Template, file fs.FS) func(w http.ResponseWriter, r *http.Request) {
	currentTmpl, err := getTmpl(tmpl, "login.tmpl", file)
	if err != nil {
		log.Error("cloning template: %w", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		page := LoginPage{
			cssPathGetter: cssPathGetter,
			recaptcha:     recaptcha,
			log:           log,
			tmpl:          currentTmpl,
		}

		switch r.Method {
		case http.MethodGet:
			page.loginGet(w, r)
		case http.MethodPost:
			page.loginAction(w, r)
		}
	}
}

func (l *LoginPage) loginGet(w http.ResponseWriter, r *http.Request) {
	model := createModel("Prihlásenie", "Login", "/login", "activities", l.cssPathGetter)
	model["csrfTokenField"] = nosurf.FormFieldName
	model["csrfToken"] = nosurf.Token(r)
	model["username"] = ""
	model["recaptchaKey"] = l.recaptcha.Key()
	if err := l.tmpl.ExecuteTemplate(w, "page", model); err != nil {
		l.log.Error("page executing context", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
	}
}

func (l *LoginPage) loginAction(w http.ResponseWriter, r *http.Request) {
	model := createModel("Prihlásenie", "Login", "/login", "activities", l.cssPathGetter)
	r.ParseForm()
	username := r.Form.Get("username")
	model["csrfToken"] = nosurf.Token(r)
	model["csrfTokenField"] = nosurf.FormFieldName
	model["username"] = username
	model["recaptchaKey"] = l.recaptcha.Key()
	err := l.recaptcha.ValidateCaptcha(r)
	if err != nil {
		l.log.Info("recaptcha validation failed", slog.Any("error", err))
		model["errorMsg"] = "Chyba pri overovaní reCAPTCHA"
	}
	if err := l.tmpl.ExecuteTemplate(w, "page", model); err != nil {
		l.log.Error("page executing context", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
	}
}
