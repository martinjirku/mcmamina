package main

import (
	"net/url"

	"jirku.sk/mcmamina/pkg/text"
)

type pageParameters struct {
	Name  string
	Path  string
	Title string
}

func (p pageParameters) PageTemplPath() string {
	path, _ := url.JoinPath("./template/pages", p.Name+".templ")
	return path
}
func (p pageParameters) PageTsPath() string {
	path, _ := url.JoinPath("./template/pages", p.Name+".ts")
	return path
}
func (p pageParameters) PageCSSName() string {
	return p.Name + ".css"
}
func (p pageParameters) PageCSSPath() string {
	path, _ := url.JoinPath("./template/pages", p.PageCSSName())
	return path
}
func (p pageParameters) ClassName() string {
	return text.CamelToDashDelimited(p.Name)
}
func (p pageParameters) PageHandlerPath() string {
	path, _ := url.JoinPath("./handlers/", p.Name+".go")
	return path
}
