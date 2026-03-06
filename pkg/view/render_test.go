package view

import (
	"net/http/httptest"
	"strings"
	"testing"

	g "maragu.dev/gomponents"
)

func TestRenderHTML_WritesDOCTYPEAndBody(t *testing.T) {
	rr := httptest.NewRecorder()
	err := RenderHTML(rr, g.Text("ok"))
	if err != nil {
		t.Fatalf("RenderHTML err: %v", err)
	}
	out := rr.Body.String()
	if !strings.Contains(out, "<!doctype html>") {
		t.Fatalf("missing doctype")
	}
	if !strings.Contains(out, "ok") {
		t.Fatalf("missing body")
	}
}
