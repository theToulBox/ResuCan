package controllers

import "github.com/llcranmer/9-2-5-Resume-Scan/views"

// Static contains the webpages
// that are not rendering dynamic content
type Static struct {
	Home    *views.View
	Privacy *views.View
	Terms   *views.View
}

// NewStatic returns a the Static webpages for
// the web app.
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("css", "static/home"),
		Privacy: views.NewView("css", "static/privacy-policy"),
		Terms:   views.NewView("css", "static/terms-and-conditions"),
	}
}
