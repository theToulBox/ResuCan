package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/llcranmer/9-2-5-Resume-Scan/controllers"
)

func main() {
	staticC := controllers.NewStatic()
	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	http.ListenAndServe(":5000", r)
}
