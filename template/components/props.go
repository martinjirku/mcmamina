package components

import (
	"github.com/a-h/templ"
)

type pageDto struct {
	ID      string
	title   string
	Content templ.Component
	Css     string
}

func NewPage(id, title, css string, body templ.Component) pageDto {
	return pageDto{
		ID:      id,
		title:   title,
		Content: body,
		Css:     css,
	}
}
