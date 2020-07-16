package controllers

import (
	"firstimeLanguage/views"
)

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "layouts/home"),
		Faq:     views.NewView("bootstrap", "layouts/faq"),
		Contact: views.NewView("bootstrap", "layouts/contact"),
	}
}

type Static struct {
	Home    *views.View
	Faq     *views.View
	Contact *views.View
}
