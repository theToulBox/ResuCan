package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/llcranmer/9-2-5-Resume-Scan/controllers"
)

func main() {
	staticC := controllers.NewStatic()
	reviewC := controllers.NewReview()

	staticDir := "/public/"
	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/privacy-policy", staticC.Privacy).Methods("GET")
	r.Handle("/terms-and-conditions", staticC.Terms).Methods("GET")

	// Resume reviews
	r.HandleFunc("/review", reviewC.New).Methods("GET")
	r.HandleFunc("/review", reviewC.Review).Methods("POST")

	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	log.Fatal(http.ListenAndServe(":5000", r))
}
