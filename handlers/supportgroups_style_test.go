package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSupportGroupsPageUsesMarketplaceStyling(t *testing.T) {
	t.Parallel()

	router := newSnapshotRouter(t)
	req := httptest.NewRequest(http.MethodGet, "/aktivity/podporne-skupiny", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.Code)
	}

	body := res.Body.String()
	for _, want := range []string{
		` text-base bg-indigo-100 py-10 px-5 md:py-10 xl:py-16 mb-0`,
		`class="h-px border-t-0 bg-gradient-to-r from-indigo-500 to-pink-500 my-2"`,
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("expected body to contain %q", want)
		}
	}

	for _, notWant := range []string{
		`bg-teal-50`,
		`text-teal-950`,
		`from-teal-100 to-cyan-500`,
	} {
		if strings.Contains(body, notWant) {
			t.Fatalf("expected body not to contain %q", notWant)
		}
	}
}
