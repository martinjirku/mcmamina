package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBurzyPageStillShowsMarketplaceContent(t *testing.T) {
	t.Parallel()

	router := newSnapshotRouter(t)
	req := httptest.NewRequest(http.MethodGet, "/aktivity/burzy", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.Code)
	}

	body := res.Body.String()
	for _, want := range []string{"Burzy", "Jarná a jesenná burza", "Handmade zimná burza"} {
		if !strings.Contains(body, want) {
			t.Fatalf("expected body to contain %q", want)
		}
	}

	for _, notWant := range []string{"Krúžky a cvičenia", "HAPPY GYM"} {
		if strings.Contains(body, notWant) {
			t.Fatalf("expected body not to contain %q", notWant)
		}
	}
}
