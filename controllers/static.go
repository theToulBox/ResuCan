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
	ResumeLength       *views.View
	HardSkills         *views.View
	SoftSkills         *views.View
}

// NewStatic returns a the Static webpages for
// the web app.
func NewStatic() *Static {
	return &Static{
		Home:               views.NewView("bootstrap", "static/home"),
		Privacy:            views.NewView("bootstrap", "static/privacy-policy"),
		Terms:              views.NewView("bootstrap", "static/terms-and-conditions"),
		About:              views.NewView("bootstrap", "static/about"),
		WhereToAddLinkedIn: views.NewView("bootstrap", "static/where-to-add-linkedin-on-resume"),
		MeasurablesExample: views.NewView("bootstrap", "static/measurables-on-resume-examples"),
		ResumeLength:       views.NewView("bootstrap", "static/why-resume-length-matters"),
		HardSkills:         views.NewView("bootstrap", "static/what-are-hard-skills"),
		SoftSkills:         views.NewView("bootstrap", "static/what-are-soft-skills"),
	}
}
