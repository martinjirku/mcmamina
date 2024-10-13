package middleware

import (
	"net/http"
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
)

var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

var httpRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Number of HTTP requests processed, labeled by status code.",
}, []string{"code", "method"})

var memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "memory_usage_bytes",
	Help: "Current memory usage in bytes.",
})

func RegisterMetrics() {
	prometheus.MustRegister(httpDuration)
	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(memoryUsage)
}

func recordMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memoryUsage.Set(float64(m.Alloc))
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := r.URL.Path // Assume you have a way to get the route
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(route))
		defer timer.ObserveDuration()

		// Wrap the ResponseWriter to capture the status code
		wrappedWriter := &responseWriter{ResponseWriter: w, statusCode: 200}
		next.ServeHTTP(wrappedWriter, r)

		// Record the status code and method
		httpRequests.WithLabelValues(http.StatusText(wrappedWriter.statusCode), r.Method).Inc()

		// Optionally, update other metrics here
		recordMemoryUsage()
	})
}
