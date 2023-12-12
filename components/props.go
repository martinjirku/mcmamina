package components

import (
	"github.com/a-h/templ"
)

type FullWidthCardDto struct {
	class      string
	background string
	padding    string
	margin     string
}

func NewFullWidthCard() *FullWidthCardDto {
	return &FullWidthCardDto{
		background: "bg-indigo-100",
		padding:    "py-12 px-5 md:py-16 xl:py-20",
		margin:     "my-12 md:my-16 xl:my-28",
	}
}

func (f FullWidthCardDto) Class(class string) FullWidthCardDto {
	f.class = class
	return f
}

func (f FullWidthCardDto) Background(background string) FullWidthCardDto {
	f.background = background
	return f
}

func (f FullWidthCardDto) Padding(padding string) FullWidthCardDto {
	f.padding = padding
	return f
}

func (f FullWidthCardDto) Margin(margin string) FullWidthCardDto {
	f.margin = margin
	return f
}

type pageDto struct {
	ID      string
	title   string
	Content templ.Component
}

func NewPage(id, title string, body templ.Component) pageDto {
	return pageDto{
		ID:      id,
		title:   title,
		Content: body,
	}
}
