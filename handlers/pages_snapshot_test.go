package handlers_test

import (
	"context"
	"errors"
	"html/template"
	"io"
	"io/fs"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/42atomys/sprout"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"jirku.sk/mcmamina/handlers"
	"jirku.sk/mcmamina/pkg/models"
	"jirku.sk/mcmamina/pkg/testutil"
)

// Snapshot usage:
//   - verify mode (default): go test ./...
//   - update mode: UPDATE_SNAPSHOTS=1 go test ./...
//
// Update snapshots only when intentional UI/template changes are made.
func TestPageSnapshots(t *testing.T) {
	t.Parallel()

	router := newSnapshotRouter(t)
	tests := []struct {
		name         string
		method       string
		route        string
		wantStatus   int
		snapshotFile string
		body         url.Values
		skipReason   string
	}{
		{
			name:       "homepage",
			method:     http.MethodGet,
			route:      "/",
			wantStatus: http.StatusOK,
			skipReason: "TODO: homepage uses time.Now() directly for two-week calendar; inject a controllable clock for deterministic snapshots.",
		},
		{name: "error", method: http.MethodGet, route: "/error", wantStatus: http.StatusOK, snapshotFile: "GET__error.snap.html"},
		{name: "about-us", method: http.MethodGet, route: "/o-nas", wantStatus: http.StatusOK, snapshotFile: "GET__o-nas.snap.html"},
		{name: "activities", method: http.MethodGet, route: "/aktivity", wantStatus: http.StatusOK, snapshotFile: "GET__aktivity.snap.html"},
		{name: "activity-baby-delivery", method: http.MethodGet, route: "/aktivity/predporodny-kurz", wantStatus: http.StatusOK, snapshotFile: "GET__aktivity__predporodny-kurz.snap.html"},
		{name: "activity-support-groups", method: http.MethodGet, route: "/aktivity/podporne-skupiny", wantStatus: http.StatusOK, snapshotFile: "GET__aktivity__podporne-skupiny.snap.html"},
		{name: "activity-marketplace", method: http.MethodGet, route: "/aktivity/burzy", wantStatus: http.StatusOK, snapshotFile: "GET__aktivity__burzy.snap.html"},
		{name: "activity-calendar", method: http.MethodGet, route: "/aktivity/kalendar", wantStatus: http.StatusOK, snapshotFile: "GET__aktivity__kalendar.snap.html"},
		{name: "supported-us", method: http.MethodGet, route: "/podpora", wantStatus: http.StatusOK, snapshotFile: "GET__podpora.snap.html"},
		{name: "tax-bonus", method: http.MethodGet, route: "/podpora/2-percenta-z-dane", wantStatus: http.StatusOK, snapshotFile: "GET__podpora__2-percenta-z-dane.snap.html"},
		{name: "volunteers", method: http.MethodGet, route: "/podpora/dobrovolnici", wantStatus: http.StatusOK, snapshotFile: "GET__podpora__dobrovolnici.snap.html"},
		{name: "login-get", method: http.MethodGet, route: "/prihlasenie", wantStatus: http.StatusOK, snapshotFile: "GET__prihlasenie.snap.html"},
		{name: "login-post-invalid", method: http.MethodPost, route: "/prihlasenie", wantStatus: http.StatusOK, snapshotFile: "POST__prihlasenie.snap.html", body: url.Values{"username": {"snapshot-user@example.com"}}},
		{
			name:       "admin",
			method:     http.MethodGet,
			route:      "/admin",
			wantStatus: http.StatusOK,
			skipReason: "TODO: /admin is behind auth middleware in production; add an authenticated test fixture and session setup.",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.skipReason != "" {
				t.Skip(tc.skipReason)
			}

			var body io.Reader
			if tc.body != nil {
				body = strings.NewReader(tc.body.Encode())
			}
			req := httptest.NewRequest(tc.method, tc.route, body)
			if tc.body != nil {
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			}

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			if res.Code != tc.wantStatus {
				t.Fatalf("route %s: expected status %d, got %d", tc.route, tc.wantStatus, res.Code)
			}

			snapshotPath := filepath.Join("testdata", "snapshots", "pages", tc.snapshotFile)
			testutil.AssertSnapshot(t, tc.route, snapshotPath, res.Body.String())
		})
	}
}

func newSnapshotRouter(t *testing.T) http.Handler {
	t.Helper()

	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	rootFS := os.DirFS("..")
	tmpl := parseTemplates(t, rootFS)

	css := stubCSS{}
	events := stubEvents{}
	sponsors := stubSponsors{}
	recaptcha := stubRecaptcha{}

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.NewPageHandler(logger, events, css, tmpl, rootFS).ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/error", handlers.HandleError(logger, css, tmpl, rootFS)).Methods(http.MethodGet)
	router.HandleFunc("/o-nas", handlers.AboutUs(logger, css, tmpl, rootFS)).Methods(http.MethodGet)
	router.HandleFunc("/aktivity", handlers.Activities(logger, css, tmpl, rootFS)).Methods(http.MethodGet)
	router.HandleFunc("/aktivity/predporodny-kurz", handlers.BabyDeliveryCourse(logger, css, tmpl, rootFS)).Methods(http.MethodGet)
	router.HandleFunc("/aktivity/podporne-skupiny", handlers.SupportGroups(logger, css, tmpl, rootFS)).Methods(http.MethodGet)
	router.HandleFunc("/aktivity/burzy", handlers.Marketplace(logger, css, tmpl, rootFS)).Methods(http.MethodGet)
	router.HandleFunc("/aktivity/kalendar", handlers.Calendar(logger, css, tmpl, rootFS)).Methods(http.MethodGet)
	router.HandleFunc("/podpora", handlers.SupportedUs(logger, css, sponsors, tmpl, rootFS)).Methods(http.MethodGet)
	router.HandleFunc("/podpora/2-percenta-z-dane", handlers.TaxBonus(logger, css, tmpl, rootFS)).Methods(http.MethodGet)
	router.HandleFunc("/podpora/dobrovolnici", handlers.Volunteers(logger, css, tmpl, rootFS)).Methods(http.MethodGet)
	router.HandleFunc("/prihlasenie", handlers.Login(logger, css, recaptcha, tmpl, oauth2.Config{}, rootFS)).Methods(http.MethodGet, http.MethodPost)

	admin := handlers.AdminHandlers{
		CssPathGetter: css,
		Recaptcha:     recaptcha,
		Log:           logger,
	}
	admin.InitTmpl(tmpl, rootFS)
	router.HandleFunc("/admin", admin.DashboardGet).Methods(http.MethodGet)

	return router
}

func parseTemplates(t *testing.T, rootFS fs.FS) *template.Template {
	t.Helper()

	tmpl, err := template.New("base").Funcs(sprout.FuncMap()).ParseFS(rootFS, "templates/**/*.tmpl")
	if err != nil {
		t.Fatalf("parse templates: %v", err)
	}
	return tmpl
}

type stubCSS struct{}

func (stubCSS) GetCssPath() (string, error) {
	return "/assets/test.css", nil
}

type stubEvents struct{}

func (stubEvents) GetEvents(_ context.Context, _, _ time.Time) ([]models.Event, error) {
	return nil, nil
}

type stubSponsors struct{}

func (stubSponsors) GetSponsors(_ context.Context) ([]models.Sponsor, error) {
	return []models.Sponsor{
		{Img: "sponsor-a.png", Url: "https://example.com/a"},
		{Img: "sponsor-b.png", Url: "https://example.com/b"},
	}, nil
}

type stubRecaptcha struct{}

func (stubRecaptcha) ValidateCaptcha(_ *http.Request) error {
	return errors.New("forced test captcha failure")
}

func (stubRecaptcha) Key() string {
	return "recaptcha-test-key"
}
