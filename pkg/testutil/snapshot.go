package testutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const updateSnapshotsEnv = "UPDATE_SNAPSHOTS"

// AssertSnapshot compares actual HTML against a golden file.
//
// Verify mode (default): go test ./...
// Update mode: UPDATE_SNAPSHOTS=1 go test ./...
func AssertSnapshot(t *testing.T, route string, snapshotPath string, actualHTML string) {
	t.Helper()

	normalizedActual := normalizeHTML(actualHTML)
	update := os.Getenv(updateSnapshotsEnv) == "1"

	if update {
		if err := os.MkdirAll(filepath.Dir(snapshotPath), 0o755); err != nil {
			t.Fatalf("route %s: create snapshot dir for %s: %v", route, snapshotPath, err)
		}
		if err := os.WriteFile(snapshotPath, []byte(normalizedActual+"\n"), 0o644); err != nil {
			t.Fatalf("route %s: write snapshot %s: %v", route, snapshotPath, err)
		}
		return
	}

	expectedBytes, err := os.ReadFile(snapshotPath)
	if err != nil {
		t.Fatalf("route %s: missing snapshot %s (%v). Regenerate with UPDATE_SNAPSHOTS=1 go test ./...", route, snapshotPath, err)
	}

	normalizedExpected := normalizeHTML(string(expectedBytes))
	if normalizedExpected == normalizedActual {
		return
	}

	idx := mismatchIndex(normalizedExpected, normalizedActual)
	expectedSnippet := snippetAround(normalizedExpected, idx, 120)
	actualSnippet := snippetAround(normalizedActual, idx, 120)

	t.Fatalf(
		"snapshot mismatch\nroute: %s\nsnapshot: %s\nfirst_diff_index: %d\nexpected_snippet: %q\nactual_snippet:   %q\n",
		route,
		snapshotPath,
		idx,
		expectedSnippet,
		actualSnippet,
	)
}

func normalizeHTML(input string) string {
	normalized := strings.ReplaceAll(input, "\r\n", "\n")
	return strings.TrimSpace(normalized)
}

func mismatchIndex(a, b string) int {
	max := len(a)
	if len(b) < max {
		max = len(b)
	}
	for i := 0; i < max; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	if len(a) != len(b) {
		return max
	}
	return -1
}

func snippetAround(s string, idx, context int) string {
	if idx < 0 {
		idx = 0
	}
	start := idx - context
	if start < 0 {
		start = 0
	}
	end := idx + context
	if end > len(s) {
		end = len(s)
	}
	return fmt.Sprintf("...%s...", s[start:end])
}
