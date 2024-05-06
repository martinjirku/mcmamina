package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/justinas/nosurf"
	"golang.org/x/oauth2"
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
	config        oauth2.Config
}

func Login(log *slog.Logger, cssPathGetter CSSPathGetter, recaptcha RecaptchaValidator, tmpl *template.Template, config oauth2.Config, file fs.FS) func(w http.ResponseWriter, r *http.Request) {
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
			config:        config,
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
	model := createModel("Prihlásenie", "/login", "activities", l.cssPathGetter)
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
	model := createModel("Prihlásenie", "/login", "activities", l.cssPathGetter)
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
	} else if username != "martinjirku@gmail.com" {
		model["errorMsg"] = "Nesprávna emailová adresa"
	} else {
		url := l.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		l.log.Info("redirection to", slog.Any("string", url))
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}
	if err := l.tmpl.ExecuteTemplate(w, "page", model); err != nil {
		l.log.Error("page executing context", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
	}
}

func GoogleCallbackHandler(config oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		code := r.FormValue("code")
		token, err := config.Exchange(ctx, code)
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Get user info
		client := config.Client(ctx, token)
		userInfoResp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if err != nil {
			http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer userInfoResp.Body.Close()
		userInfoBytes, err := io.ReadAll(userInfoResp.Body)
		if err != nil {
			http.Error(w, "Failed to read response: "+err.Error(), http.StatusInternalServerError)
			return
		}

		var user struct {
			Name          string
			ID            string
			Email         string
			VerifiedEmail string
			Picture       string
		}
		if err := json.Unmarshal(userInfoBytes, &user); err != nil {
			http.Error(w, "Failed to parse user info: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Render a response containing user info
		fmt.Fprintf(w, `<html><head><title>Google User Info</title></head>
        <body><h1>Welcome %s!</h1>
			<p><strong>ID:</strong> %s</p><p><strong>Email:</strong> %s</p><p><strong>Email Verified:</strong> %v</p>
			<p><strong>Profile Picture:</strong> <img src="%s" alt="Profile Picture"></p>
		</body></html>`,
			user.Name, user.ID, user.Email, user.VerifiedEmail, user.Picture)

	}
}
