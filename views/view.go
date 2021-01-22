package views

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/csrf"
)

var (
	// LayoutDir is the directory containing
	// parts of the website that define
	// the whole
	LayoutDir string = "views/layouts/"
	// TemplateDir contains all of the templates to be rendered
	// as pages
	TemplateDir string = "views/"
	// TemplateExt is the ending of the template files
	TemplateExt string = ".gohtml"
)

// View contains data that is used to render
// webpage
type View struct {
	Template *template.Template
	Layout   string
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}

// NewView is used to parse all of the templates
// and return a view
func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles()...)
	t, err := template.New("").Funcs(template.FuncMap{
		"csrfField": func() (template.HTML, error) {
			return "", errors.New("csrfField is not implemented")
		},
	}).ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// Render accepts a ptr to a View and then executes
// the template(s), filling in the data, to build
// webpages
//Render is used to render the view with the predefined layout.
func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	var vd Data
	switch d := data.(type) {
	case Data:
		vd = d
		//do nothing
	default:
		vd = Data{
			Yield: data,
		}
	}
	var buf bytes.Buffer
	csrfField := csrf.TemplateField(r)
	tpl := v.Template.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrfField
		},
	})
	if err := tpl.ExecuteTemplate(&buf, v.Layout, vd); err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong, if the problem persists, please email support@lenslocked.com", http.StatusInternalServerError)
		return
	}
	_, _ = io.Copy(w, &buf)
}

// ServeHTTP accepts a ptr to a view and renders
// the view view writing it as http.ResponseWriter
func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, r, nil)
}
