package controllers

import (
	"firstimeLanguage.com/views"
)

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Faq:     views.NewView("bootstrap", "static/faq"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}

type Static struct {
	Home    *views.View
	Faq     *views.View
	Contact *views.View
}
