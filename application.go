package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/llcranmer/9-2-5-Resume-Scan/controllers"
)

func main() {
	staticC := controllers.NewStatic()
	r := mux.NewRouter()
	r.Handle("/home", staticC.Home).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
