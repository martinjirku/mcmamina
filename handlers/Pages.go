package handlers

import (
	"fmt"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
)

func createModel(title, menuItemUrl, class string, cssPathGetter CSSPathGetter) map[string]any {
	cssPath, _ := cssPathGetter.GetCssPath()
	return map[string]any{
		"Title":       title,
		"Css":         cssPath,
		"layoutClass": fmt.Sprintf("%s w-full bg-cover bg-center text-indigo-800 font-light", class),
		"Menu":        getMenuItems(menuItemUrl),
	}
}

func addActivitySubmenu(model map[string]any, activePath string) map[string]any {
	model["submenu"] = []map[string]any{
		{"label": "Predpôrodný kurz", "href": "/aktivity/predporodny-kurz", "isActive": activePath == "/aktivity/predporodny-kurz"},
		{"label": "Podporné skupiny", "href": "/aktivity/podporne-skupiny", "isActive": activePath == "/aktivity/podporne-skupiny"},
		{"label": "Burzy", "href": "/aktivity/burzy", "isActive": activePath == "/aktivity/burzy"},
		{"label": "Kalendár", "href": "/aktivity/kalendar", "isActive": activePath == "/aktivity/kalendar"},
	}
	return model
}

func getTmpl(tmpl *template.Template, templateName string, file fs.FS) (*template.Template, error) {
	currentTmpl, err := tmpl.Clone()
	if err != nil {
		return nil, fmt.Errorf("cloning template: %w", err)
	}
	currentTmpl, err = currentTmpl.ParseFS(file, fmt.Sprintf("templates/pages/%s", templateName))
	if err != nil {
		return nil, fmt.Errorf("ParseFS 'templates/pages/%s': %w", templateName, err)
	}
	return currentTmpl, nil
}

func AboutUs(log *slog.Logger, cssPathGetter CSSPathGetter, tmpl *template.Template, file fs.FS) func(w http.ResponseWriter, r *http.Request) {
	currentTmpl, err := getTmpl(tmpl, "about-us.tmpl", file)
	if err != nil {
		log.Error("cloning template: %w", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		model := createModel("O nás", "/o-nas", "about-us", cssPathGetter)
		if err := currentTmpl.ExecuteTemplate(w, "page", model); err != nil {
			log.Error("page executing context", err)
			http.Redirect(w, r, "/error", http.StatusInternalServerError)
		}
	}
}

func Activities(log *slog.Logger, cssPathGetter CSSPathGetter, tmpl *template.Template, file fs.FS) func(w http.ResponseWriter, r *http.Request) {
	currentTmpl, err := getTmpl(tmpl, "activity.tmpl", file)
	if err != nil {
		log.Error("cloning template: %w", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		model := createModel("Aktivity", "/aktivity", "activities", cssPathGetter)
		model = addActivitySubmenu(model, "/aktivity")
		if err := currentTmpl.ExecuteTemplate(w, "page", model); err != nil {
			log.Error("page executing context", err)
			http.Redirect(w, r, "/error", http.StatusInternalServerError)
		}
	}
}

func BabyDeliveryCourse(log *slog.Logger, cssPathGetter CSSPathGetter, tmpl *template.Template, file fs.FS) func(w http.ResponseWriter, r *http.Request) {
	currentTmpl, err := getTmpl(tmpl, "activity.babydelivery.tmpl", file)
	if err != nil {
		log.Error("cloning template: %w", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		model := createModel("Predpôrodný kurz", "/aktivity", "baby-delivery-course", cssPathGetter)
		model = addActivitySubmenu(model, "/aktivity/predporodny-kurz")
		if err := currentTmpl.ExecuteTemplate(w, "page", model); err != nil {
			log.Error("page executing context", err)
			http.Redirect(w, r, "/error", http.StatusInternalServerError)
		}
	}
}

func Marketplace(log *slog.Logger, cssPathGetter CSSPathGetter, tmpl *template.Template, file fs.FS) func(w http.ResponseWriter, r *http.Request) {
	currentTmpl, err := getTmpl(tmpl, "activity.marketplace.tmpl", file)
	if err != nil {
		log.Error("cloning template: %w", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		model := createModel("Burzy", "/aktivity", "marketplace", cssPathGetter)
		model = addActivitySubmenu(model, "/aktivity/burzy")
		if err := currentTmpl.ExecuteTemplate(w, "page", model); err != nil {
			log.Error("page executing context", err)
			http.Redirect(w, r, "/error", http.StatusInternalServerError)
		}
	}
}

func SupportGroups(log *slog.Logger, cssPathGetter CSSPathGetter, tmpl *template.Template, file fs.FS) func(w http.ResponseWriter, r *http.Request) {
	currentTmpl, err := getTmpl(tmpl, "activity.marketplace.tmpl", file)
	if err != nil {
		log.Error("cloning template: %w", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		model := createModel("Podporné skupiny", "/aktivity", "activities", cssPathGetter)
		model = addActivitySubmenu(model, "/aktivity/podporne-skupiny")
		if err := currentTmpl.ExecuteTemplate(w, "page", model); err != nil {
			log.Error("page executing context", err)
			http.Redirect(w, r, "/error", http.StatusInternalServerError)
		}
	}
}

func Calendar(log *slog.Logger, cssPathGetter CSSPathGetter, tmpl *template.Template, file fs.FS) func(w http.ResponseWriter, r *http.Request) {
	currentTmpl, err := getTmpl(tmpl, "activity.calendar.tmpl", file)
	if err != nil {
		log.Error("cloning template: %w", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		model := createModel("Kalendár", "/aktivity", "calendar", cssPathGetter)
		model = addActivitySubmenu(model, "/aktivity/kalendar")
		if err := currentTmpl.ExecuteTemplate(w, "page", model); err != nil {
			log.Error("page executing context", err)
			http.Redirect(w, r, "/error", http.StatusInternalServerError)
		}
	}
}
