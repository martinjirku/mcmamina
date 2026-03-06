package components

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

func TestFullWidth_UsesDefaults(t *testing.T) {
	out := renderNode(t, FullWidth(FullWidthProps{}, g.Text("x")))
	if !strings.Contains(out, "w-full bg-opacity-95 text-xl leading-10 flex justify-around") {
		t.Fatalf("missing base classes: %s", out)
	}
	if !strings.Contains(out, "bg-indigo-100") {
		t.Fatalf("missing default bg: %s", out)
	}
	if !strings.Contains(out, "py-12 px-5 md:py-16 xl:py-20") {
		t.Fatalf("missing default padding: %s", out)
	}
	if !strings.Contains(out, "my-12 md:my-16 xl:my-28") {
		t.Fatalf("missing default margins: %s", out)
	}
}

func TestCardContent_AppendsClass(t *testing.T) {
	out := renderNode(t, CardContent("text-red-500", g.Text("x")))
	if !strings.Contains(out, "w-full md:max-w-3xl lg:max-w-4xl text-red-500") {
		t.Fatalf("unexpected card classes: %s", out)
	}
}

func TestHomeIcon_HasViewBoxAndClass(t *testing.T) {
	out := renderNode(t, HomeIcon("fill-black", "16"))
	if !strings.Contains(out, `viewBox="0 -960 960 960"`) {
		t.Fatalf("missing viewBox: %s", out)
	}
	if !strings.Contains(out, `class="inline-block fill-black"`) {
		t.Fatalf("missing class passthrough: %s", out)
	}
}

func TestAnimatedLogo_HasExpectedClasses(t *testing.T) {
	out := renderNode(t, AnimatedLogo("w-48"))
	if !strings.Contains(out, "mcmamina-logo mcmamina-animate w-48") {
		t.Fatalf("missing logo classes: %s", out)
	}
	if !strings.Contains(out, `id="heart"`) {
		t.Fatalf("missing heart path: %s", out)
	}
}
