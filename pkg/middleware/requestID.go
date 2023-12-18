package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/segmentio/ksuid"
)

type key int

const RequestIDKey key = 0

func RequestID(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			requestID, err := ksuid.NewRandom()
			if err != nil {
				logger.Error("Failed to generate request ID", slog.Any("error", err))
			} else {
				id := requestID.String()
				w.Header().Set("X-Request-ID", id)
				logger.Debug("RequestID assigned", slog.String("request-id", id))
				ctx = context.WithValue(r.Context(), RequestIDKey, id)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
