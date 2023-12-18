package middleware

import (
	"log/slog"
	"net/http"
)

func Logger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("referer", r.Referer()),
				slog.String("user-agent", r.UserAgent()),
			)
			next.ServeHTTP(w, r)
		})
	}
}
