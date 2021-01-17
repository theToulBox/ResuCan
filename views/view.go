package views

import (
	"html/template"
	"net/http"
	"path/filepath"
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
	t, err := template.ParseFiles(files...)
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
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)

}

// ServeHTTP accepts a ptr to a view and renders
// the view view writing it as http.ResponseWriter
func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)
}
