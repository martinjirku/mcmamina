package handlers

import (
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/justinas/nosurf"
	"jirku.sk/mcmamina/pkg/middleware"
)

type AdminHandlers struct {
	CssPathGetter CSSPathGetter
	Recaptcha     RecaptchaValidator
	Log           *slog.Logger
	loginTmpl     *template.Template
}

func (h *AdminHandlers) InitTmpl(tmpl *template.Template, file fs.FS) *AdminHandlers {
	var err error
	h.loginTmpl, err = getTmpl(tmpl, "admin.tmpl", file)
	if err != nil {
		h.Log.Error("cloning template: %w", err)
	}
	return h
}

func (h *AdminHandlers) DashboardGet(w http.ResponseWriter, r *http.Request) {
	model := createModel("Prihlásenie", "/login", "activities", h.CssPathGetter)
	model["csrfTokenField"] = nosurf.FormFieldName
	model["csrfToken"] = nosurf.Token(r)
	model["user"] = middleware.GetUser(r)
	model["recaptchaKey"] = h.Recaptcha.Key()

	if err := h.loginTmpl.ExecuteTemplate(w, "page", model); err != nil {
		h.Log.Error("page executing context", err)
		http.Redirect(w, r, "/error", http.StatusInternalServerError)
	}
}
