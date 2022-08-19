package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/example/{name}", ExampleHandler)

	http.ListenAndServe(":8000", r)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	log.Printf("name is: %v", vars["name"])
}
