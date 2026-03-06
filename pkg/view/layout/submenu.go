package layout

import (
	"strings"

	"jirku.sk/mcmamina/pkg/view/components"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func Submenu(items []MenuItem) g.Node {
	nodes := make([]g.Node, 0, len(items))
	for _, item := range items {
		divClass := "my-1 align-middle text-center overflow-hidden text-ellipsis whitespace-nowrap"
		if item.IsActive {
			divClass += " underline underline-offset-2"
		}
		nodes = append(nodes,
			h.Div(
				h.Class(divClass),
				h.A(
					h.Class(strings.TrimSpace("inline-block py-2 px-2 sm:px-3 text-xs md:text-sm hover:bg-slate-800 rounded-md transition-colors duration-500 whitespace-nowrap text-ellipsis")),
					h.Href(item.Href),
					g.Text(item.Label),
				),
			),
		)
	}

	return components.FullWidth(
		components.FullWidthProps{M: "m-0", P: "p-0", Bg: "bg-slate-700", Class: "text-slate-100"},
		components.CardContent("flex flex-row justify-center gap-2 md:gap-6 p-1", g.Group(nodes)),
	)
}
