package handlers

import (
	"log/slog"
	"net/http"

	"jirku.sk/mcmamina/template/components"
	"jirku.sk/mcmamina/template/pages"
)

func AboutUs(Log *slog.Logger, cssPathGetter CSSPathGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		cssPath, _ := cssPathGetter.GetCssPath()
		// TODO: handle GET
		components.Page(components.NewPage(
			"AboutUs",
			"Domov",
			cssPath,
			pages.AboutUsPage(),
		)).Render(r.Context(), w)
	}
}

func Activities(Log *slog.Logger, cssPathGetter CSSPathGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		cssPath, _ := cssPathGetter.GetCssPath()
		// TODO: handle GET
		components.Page(components.NewPage(
			"Activities",
			"Aktivity",
			cssPath,
			pages.ActivitiesPage(),
		)).Render(r.Context(), w)
	}
}

func BabyDeliveryCourse(Log *slog.Logger, cssPathGetter CSSPathGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		cssPath, _ := cssPathGetter.GetCssPath()
		// TODO: handle GET
		components.Page(components.NewPage(
			"BabyDeliveryCourse",
			"Predpôrodný kurz",
			cssPath,
			pages.BabyDeliveryCoursePage(),
		)).Render(r.Context(), w)
	}
}

func Calendar(Log *slog.Logger, cssPathGetter CSSPathGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		cssPath, _ := cssPathGetter.GetCssPath()
		// TODO: handle GET
		components.Page(components.NewPage(
			"Calendar",
			"Kalendár",
			cssPath,
			pages.CalendarPage(),
		)).Render(r.Context(), w)
	}
}

func Marketplace(Log *slog.Logger, cssPathGetter CSSPathGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		cssPath, _ := cssPathGetter.GetCssPath()
		// TODO: handle GET
		components.Page(components.NewPage(
			"Marketplace",
			"Burzy",
			cssPath,
			pages.MarketplacePage(),
		)).Render(r.Context(), w)
	}
}

func SupportGroups(Log *slog.Logger, cssPathGetter CSSPathGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		cssPath, _ := cssPathGetter.GetCssPath()
		// TODO: handle GET
		components.Page(components.NewPage(
			"SupportGroups",
			"Podporné skupiny",
			cssPath,
			pages.SupportGroupsPage(),
		)).Render(r.Context(), w)
	}
}

func TaxBonus(Log *slog.Logger, cssPathGetter CSSPathGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		cssPath, _ := cssPathGetter.GetCssPath()
		// TODO: handle GET
		components.Page(components.NewPage(
			"TaxBonus",
			"2 Percentá z dane",
			cssPath,
			pages.TaxBonusPage(),
		)).Render(r.Context(), w)
	}
}

func Volunteers(Log *slog.Logger, cssPathGetter CSSPathGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		cssPath, _ := cssPathGetter.GetCssPath()
		// TODO: handle GET
		components.Page(components.NewPage(
			"Volunteers",
			"Dobrovoľníci",
			cssPath,
			pages.VolunteersPage(),
		)).Render(r.Context(), w)
	}
}
