package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAktivityPageWeeklyProgramUsesPageStyles(t *testing.T) {
	t.Parallel()

	router := newSnapshotRouter(t)
	req := httptest.NewRequest(http.MethodGet, "/aktivity", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.Code)
	}

	body := res.Body.String()
	for _, want := range []string{
		"<h1 class=\"text-2xl font-bold pb-4\">Krúžky a cvičenia</h1>",
		"<h2 class=\"text-xl font-bold pt-2\">PONDELOK</h2>",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("expected body to contain %q", want)
		}
	}

	for _, notWant := range []string{
		"bg-teal-50",
		"text-teal-950",
		"from-teal-100 to-cyan-500",
	} {
		if strings.Contains(body, notWant) {
			t.Fatalf("expected body not to contain %q", notWant)
		}
	}
}
