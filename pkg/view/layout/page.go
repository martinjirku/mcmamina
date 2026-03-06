package layout

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type PageProps struct {
	Title  string
	Css    string
	Module string
}

func PageDocument(props PageProps, content g.Node) g.Node {
	headChildren := []g.Node{
		h.Meta(g.Attr("charSet", "utf-8")),
		h.Meta(h.Name("viewport"), h.Content("width=device-width,initial-scale=1")),
		h.TitleEl(g.Text(props.Title)),
	}
	if props.Css != "" {
		headChildren = append(headChildren, h.Link(h.Rel("stylesheet"), h.Href(props.Css)))
	}
	headChildren = append(headChildren, h.Link(h.Rel("icon"), h.Href("favicon.ico")))

	bodyChildren := []g.Node{content}
	if props.Module != "" {
		bodyChildren = append(bodyChildren, h.Script(g.Attr("type", "module"), h.Src("/"+props.Module+".es")))
	}

	return h.HTML(g.Attr("lang", "sk"), h.Class("h-full overflow-x-clip"),
		h.Head(headChildren...),
		h.Body(append([]g.Node{h.Class("h-full")}, bodyChildren...)...),
	)
}
