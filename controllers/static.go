package controllers

import "github.com/theToulBox/ResuCan/views"

// Static contains the webpages
// that are not rendering dynamic content
type Static struct {
	Home               *views.View
	Privacy            *views.View
	Terms              *views.View
	About              *views.View
	WhereToAddLinkedIn *views.View
	MeasurablesExample *views.View
}

// NewStatic returns a the Static webpages for
// the web app.
func NewStatic() *Static {
	return &Static{
		Home:               views.NewView("css", "static/home"),
		Privacy:            views.NewView("css", "static/privacy-policy"),
		Terms:              views.NewView("css", "static/terms-and-conditions"),
		About:              views.NewView("css", "static/about"),
		WhereToAddLinkedIn: views.NewView("css", "static/where-to-add-linkedin-on-resume"),
		MeasurablesExample: views.NewView("css", "static/measurables-on-resume-examples"),
	}
}
