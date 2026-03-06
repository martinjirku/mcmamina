package layout

import (
	"fmt"
	"strings"

	"jirku.sk/mcmamina/pkg/view/components"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type MenuItem struct {
	Label    string
	Href     string
	IsActive bool
}

type LayoutProps struct {
	LayoutClass string
	Menu        []MenuItem
}

func MainLayout(props LayoutProps, children ...g.Node) g.Node {
	mainClass := strings.TrimSpace("page-background flex-grow " + props.LayoutClass)
	return h.Div(
		h.Class("h-screen flex flex-col "+props.LayoutClass),
		h.Header(
			h.Class("w-full z-50 bg-indigo-400 sticky shadow-lg top-0 text-cyan-50 text-xl"),
			h.Nav(
				h.Class("hidden md:flex justify-around py-6"),
				g.Attr("aria-label", "Hlavné"),
				h.Ul(
					h.Class("flex gap-6"),
					g.Group(renderMenuItems(props.Menu)),
				),
			),
			h.Nav(
				h.Class("mobile-menu w-full md:hidden left-0 flex-col"),
				g.Attr("aria-label", "Hlavné"),
				h.Div(
					h.Class("w-full flex relative bg-indigo-400 z-30 py-4"),
					h.H1(h.Class("font-light p-2 py-2 text-2xl"), h.A(h.Href("/"), g.Text("Mc Mamina"))),
					h.Button(h.Class("hamburger stroke-indigo-100 absolute right-0 pr-2"), g.Attr("aria-label", "Open menu"), AnimatedHamburger()),
				),
				h.Ul(
					h.Class("flex z-10 flex-col gap-3 absolute top-20 bg-indigo-400 w-full text-cyan-50 transition-transform ease-in-out duration-500 transform h-0 overflow-hidden -translate-y-full opacity-0"),
					g.Group(renderMenuItems(props.Menu)),
				),
			),
		),
		h.Main(append([]g.Node{h.Class(mainClass)}, children...)...),
		h.Script(g.Attr("type", "module"), h.Src("/layout.es")),
		Footer(),
	)
}

func renderMenuItems(items []MenuItem) []g.Node {
	out := make([]g.Node, 0, len(items))
	for _, item := range items {
		class := "rounded-md transition-colors ease-out-in delay-50 duration-200 text-indigo-50 py-4 px-4 inline-block w-full h-full bg-indigo-400 hover:bg-indigo-500 hover:opacity-90"
		if item.IsActive {
			class = "rounded-md transition-colors ease-out-in delay-50 duration-200 text-indigo-50 py-4 px-4 inline-block w-full h-full bg-indigo-500 opacity-70"
		}
		out = append(out, h.Li(
			h.Class("inline-block align-middle relative"),
			h.A(h.Class(class), h.Href(item.Href), g.Text(item.Label)),
		))
	}
	return out
}

func AnimatedHamburger() g.Node {
	return g.Raw(`<svg height="46" width="46" class="fill-none mc-animated-hamburger" viewBox="0 0 60 40">
	<g transform="matrix(1,0,0,1,-389.5,-264.004)">
		<g>
			<g transform="matrix(1,0,0,1,0,5)">
				<path id="top" class="class" d="M390,270L420,270L420,270C420,270 420.195,250.19 405,265C389.805,279.81 390,279.967 390,279.967"></path>
			</g>
			<g transform="matrix(1,1.22465e-16,1.22465e-16,-1,0.00024296,564.935)">
				<path id="bottom" class="class" d="M390,270L420,270L420,270C420,270 420.195,250.19 405,265C389.805,279.81 390,279.967 390,279.967"></path>
			</g>
			<path class="class" id="middle" d="M390,284.967L420,284.967"></path>
		</g>
	</g>
</svg>`)
}

func Footer() g.Node {
	iconClass := "fill-indigo-100 mr-2"
	iconDimension := "16"

	return h.Footer(
		h.Class("footer w-full flex pt-10 px-6 pb-16 shadow-xl content-center justify-around bg-neutral-900 h-min-40 text-indigo-100 sticky"),
		h.Div(
			h.Class("flex flex-col flex-wrap sm:flex-row gap-10 md:gap-x-28 justify-between xl:flex-nowrap md:max-w-5xl"),
			h.Address(
				h.Class("flex-grow not-italic font-thin text-sm"),
				h.H2(h.Class("font-bold leading-10 underline underline-offset-4"), g.Text("Materské centrum MAMINA o.z.")),
				h.H1(h.Class("font-bold leading-10 underline underline-offset-4"), g.Text("Fakturačné údaje")),
				h.Span(components.HomeIcon(iconClass, iconDimension), g.Text(" Tatranská 10, 97411 Banská Bystrica")),
				h.Br(),
				h.Span(g.Text("IČO: 37956825")),
				h.Br(),
				h.Span(g.Text("DIČ: 2022358239")),
				h.Br(),
				h.Span(components.BankIcon(iconClass, iconDimension), g.Text(" SK62 8330 0000 0023 0190 0933")),
				h.Br(),
			),
			h.Div(
				h.Class("flex-grow font-thin text-sm"),
				h.H1(h.Class("font-bold leading-10 underline underline-offset-4"), g.Text("Otváracie hodiny")),
				h.P(g.Text("(ported in next task)")),
			),
			h.Div(
				h.Class("flex-grow flex-shrink font-thin text-sm flex w-full md:w-auto justify-start items-center flex-col"),
				h.Div(h.Class("flex-grow-0 flex justify-around items-center pb-5"),
					h.A(h.Href("https://www.facebook.com/MaterskeCentrumMamina/"), g.Attr("aria-label", "Facebook stránka Materského centra MAMINA"), g.Attr("target", "_blank"), h.Class("p-2"), g.Attr("rel", "noreferrer"), h.TitleAttr("Facebook stránka Materského centra MAMINA"), components.FacebookIcon("fill-indigo-100 hover:animate-pulse", "42")),
					h.A(h.Href("https://www.instagram.com/mc.mamina/"), g.Attr("aria-label", "Instagram stránka Materského centra MAMINA"), g.Attr("target", "_blank"), g.Attr("rel", "noreferrer"), h.Class("p-2"), h.TitleAttr("Instagram stránka Materského centra MAMINA"), components.InstagramIcon("fill-indigo-100 hover:animate-pulse", "42")),
				),
				h.Div(h.Class("w-full flex-grow flex items-center justify-center mt-10 lg:mt-0"),
					h.A(h.Href("https://linkedin.com/in/martin-j-65786267"), g.Attr("target", "_blank"), g.Attr("rel", "noreferrer"), h.Class("px-5 py-3 bg-slate-800 rounded-md  hover:animate-spin hover:cursor-pointer"), h.Style("animation-iteration-count:1"), g.Text("Vytvoril MJ")),
				),
			),
		),
	)
}

func init() {
	_ = fmt.Sprintf
}
