package handlers

import (
	"context"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"

	"jirku.sk/mcmamina/pkg/models"
)

type SponsorGetter interface {
	GetSponsors(ctx context.Context) ([]models.Sponsor, error)
}

func addSupportSubmenu(model map[string]any, activePath string) map[string]any {
	model["submenu"] = []map[string]any{
		{"label": "Podporili nás", "href": "/podpora", "isActive": activePath == "/podpora"},
		{"label": "2% z dane", "href": "/podpora/2-percenta-z-dane", "isActive": activePath == "/podpora/2-percenta-z-dane"},
		{"label": "Dobrovoľníci", "href": "/podpora/dobrovolnici", "isActive": activePath == "/podpora/dobrovolnici"},
	}
	return model
}

func SupportedUs(log *slog.Logger, cssPathGetter CSSPathGetter, sponsorGetter SponsorGetter, tmpl *template.Template, file fs.FS) func(w http.ResponseWriter, r *http.Request) {
	currentTmpl, err := getTmpl(tmpl, "support.tmpl", file)
	if err != nil {
		log.Error("cloning template", slog.Any("error", err))
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		model := createModel("Podporili nás", "/podpora", "supported-us", cssPathGetter)
		model = addSupportSubmenu(model, "/podpora")
		model["sponsors"], _ = sponsorGetter.GetSponsors(r.Context())

		if err := currentTmpl.ExecuteTemplate(w, "page", model); err != nil {
			log.Error("page executing context", slog.Any("error", err))
			http.Redirect(w, r, "/error", http.StatusInternalServerError)
		}
	}
}

func TaxBonus(log *slog.Logger, cssPathGetter CSSPathGetter, tmpl *template.Template, file fs.FS) func(w http.ResponseWriter, r *http.Request) {
	currentTmpl, err := getTmpl(tmpl, "support.2percent.tmpl", file)
	if err != nil {
		log.Error("cloning template", slog.Any("error", err))
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		model := createModel("2 Percentá z dane", "/podpora", "tax-bonus", cssPathGetter)
		model = addSupportSubmenu(model, "/podpora/2-percenta-z-dane")
		if err := currentTmpl.ExecuteTemplate(w, "page", model); err != nil {
			log.Error("page executing context", slog.Any("error", err))
			http.Redirect(w, r, "/error", http.StatusInternalServerError)
		}
	}
}

func Volunteers(log *slog.Logger, cssPathGetter CSSPathGetter, tmpl *template.Template, file fs.FS) func(w http.ResponseWriter, r *http.Request) {
	currentTmpl, err := getTmpl(tmpl, "support.volunteers.tmpl", file)
	if err != nil {
		log.Error("cloning template", slog.Any("error", err))
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		model := createModel("Dobrovoľníci", "/podpora", "volunteers", cssPathGetter)
		model = addSupportSubmenu(model, "/podpora/dobrovolnici")
		if err := currentTmpl.ExecuteTemplate(w, "page", model); err != nil {
			log.Error("page executing context", slog.Any("error", err))
			http.Redirect(w, r, "/error", http.StatusInternalServerError)
		}
	}
}
