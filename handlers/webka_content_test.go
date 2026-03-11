package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAktivityPageContainsDetailedWeeklyProgram(t *testing.T) {
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
		"Podporné skupiny",
		"PONDELOK",
		"HAPPY GYM",
		"Deti sa najlepšie učia hrou, pohybom a spevom.",
		"Mama kruh",
		"Je bezplatný program, ktorý je určený pre znevýhodnené rodiny",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("expected body to contain %q", want)
		}
	}
}
