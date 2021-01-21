package controllers

import (
	"fmt"
	"net/http"

	"github.com/llcranmer/9-2-5-Resume-Scan/views"
)

type Review struct {
	NewView    *views.View
	ResultView *views.View
}

type ReviewForm struct {
	Resume      string `schema:"resume"`
	Description string `schema:"description"`
}

func NewReview() *Review {
	return &Review{
		NewView:    views.NewView("css", "review/new"),
		ResultView: views.NewView("css", "review/result"),
	}
}

func (re *Review) New(w http.ResponseWriter, r *http.Request) {
	re.NewView.Render(w, r, nil)
}

// Critique will analyze the Resume against the Job Description
func (re *Review) Review(w http.ResponseWriter, r *http.Request) {
	// now i have the form data container from parsing the data
	vd := views.Data{}
	form := ReviewForm{}
	if err := parseForm(r, &form); err != nil {
		vd.SetAlert(err)
		re.NewView.Render(w, r, vd)
	}

	// Run the analysis
	if err := re.Critique(form.Resume, form.Description); err != nil {
		vd.SetAlert(err)
		re.NewView.Render(w, r, vd)
	}

	// collect the res
	vd.Yield = form

	// render the view
	re.ResultView.Render(w, r, vd)
}

func (re *Review) Critique(d string, r string) error {
	fmt.Printf("R is %s, D is %s", r, d)
	return nil
}
