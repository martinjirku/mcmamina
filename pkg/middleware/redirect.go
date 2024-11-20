package middleware

import (
	"log/slog"
	"net/http"
)

func RedirectFromWWWSubdomain(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("redirecting", slog.String("host", r.Host))
			if r.Host == "www.mcmamina.sk" {
				if r.URL.RawQuery != "" {
					target := "https://mcmamina.sk"
					// target := fmt.Sprintf("https://mcmamina.sk%s", r.URL.Path)
					// target += "?" + r.URL.RawQuery
					logger.Info("redirecting", slog.String("host", r.Host), slog.String("target", target))
					http.Redirect(w, r, target, http.StatusMovedPermanently)
					return
				}

			}
			next.ServeHTTP(w, r)
		})
	}
}
