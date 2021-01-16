package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/llcranmer/9-2-5-Resume-Scan/controllers"
	"github.com/llcranmer/9-2-5-Resume-Scan/views"
)

var (
	homeView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeView.Render(w, r, nil)
}
func main() {
	staticC := controllers.NewStatic()
	r := mux.NewRouter()
	r.Handle("/home", staticC.Home).Methods("GET")
	http.ListenAndServe(":5000", nil)
}
