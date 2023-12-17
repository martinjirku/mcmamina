package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"jirku.sk/mcmamina/models"
	"jirku.sk/mcmamina/template/components"
	"jirku.sk/mcmamina/template/pages"
)

type SponsorGetter interface {
	GetSponsors(ctx context.Context) ([]models.Sponsor, error)
}

func SupportedUs(Log *slog.Logger, cssPathGetter CSSPathGetter, sponsorGetter SponsorGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info("request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		cssPath, _ := cssPathGetter.GetCssPath()
		sponsors, _ := sponsorGetter.GetSponsors(r.Context())
		// TODO: handle GET
		components.Page(components.NewPage(
			"SupportedUs",
			"Podporili nás",
			cssPath,
			pages.SupportedUsPage(sponsors),
		)).Render(r.Context(), w)
	}
}
