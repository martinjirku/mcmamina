package components

import (
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type FullWidthProps struct {
	Class string
	Bg    string
	P     string
	M     string
}

func FullWidth(props FullWidthProps, children ...g.Node) g.Node {
	bg := props.Bg
	if bg == "" {
		bg = "bg-indigo-100"
	}
	p := props.P
	if p == "" {
		p = "py-12 px-5 md:py-16 xl:py-20"
	}
	m := props.M
	if m == "" {
		m = "my-12 md:my-16 xl:my-28"
	}

	class := strings.TrimSpace(strings.Join([]string{
		"w-full bg-opacity-95 text-xl leading-10 flex justify-around",
		props.Class,
		bg,
		p,
		m,
	}, " "))

	return h.Div(append([]g.Node{h.Class(class)}, children...)...)
}

func CardContent(class string, children ...g.Node) g.Node {
	base := "w-full md:max-w-3xl lg:max-w-4xl"
	if class != "" {
		base = base + " " + class
	}
	return h.Div(append([]g.Node{h.Class(base)}, children...)...)
}
