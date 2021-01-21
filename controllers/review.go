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
	Description string `description:"description"`
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
	var vd views.Data
	var form ReviewForm
	fmt.Println("function has been hit")
	if err := parseForm(r, &form); err != nil {
		vd.SetAlert(err)
		re.NewView.Render(w, r, vd)
		return
	}
	http.Redirect(w, r, "/result", http.StatusFound)

}
