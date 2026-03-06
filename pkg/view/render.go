package view

import (
	"io"

	g "maragu.dev/gomponents"
)

func RenderHTML(w io.Writer, nodes ...g.Node) error {
	_, _ = io.WriteString(w, "<!doctype html>")
	return g.Group(nodes).Render(w)
}
