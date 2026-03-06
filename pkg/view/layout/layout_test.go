package layout

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
)

func renderNode(t *testing.T, n g.Node) string {
	t.Helper()
	var b strings.Builder
	if err := n.Render(&b); err != nil {
		t.Fatalf("render: %v", err)
	}
	return b.String()
}

func TestPageDocument_HasHeadAndCssLink(t *testing.T) {
	out := renderNode(t, PageDocument(PageProps{Title: "x", Css: "/a.css"}, g.Text("content")))
	if !strings.Contains(out, "<head>") {
		t.Fatalf("missing head: %s", out)
	}
	if !strings.Contains(out, `<link rel="stylesheet" href="/a.css">`) {
		t.Fatalf("missing css link: %s", out)
	}
}

func TestMainLayout_HasNavAndFooter(t *testing.T) {
	out := renderNode(t, MainLayout(LayoutProps{
		LayoutClass: "bg",
		Menu: []MenuItem{{Label: "Domov", Href: "/", IsActive: true}},
	}, g.Text("body")))

	if !strings.Contains(out, `<nav class="hidden md:flex justify-around py-6" aria-label="Hlavné">`) {
		t.Fatalf("missing desktop nav: %s", out)
	}
	if !strings.Contains(out, `<footer class="footer`) {
		t.Fatalf("missing footer: %s", out)
	}
}

func TestSubmenu_Highlight(t *testing.T) {
	out := renderNode(t, Submenu([]MenuItem{{Label: "A", Href: "/a", IsActive: true}}))
	if !strings.Contains(out, "underline underline-offset-2") {
		t.Fatalf("missing active highlight: %s", out)
	}
}
