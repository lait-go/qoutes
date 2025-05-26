package main

import (
	"guotes/hand"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	store := hand.NewStore()

	r := mux.NewRouter()

	r.HandleFunc("/quotes", store.QuotesHandler).Methods("GET", "POST")

	r.HandleFunc("/quotes/random", store.RandomHandler).Methods("GET")

	r.HandleFunc("/quotes/{id:[0-9]+}", store.DeleteHandler).Methods("DELETE")

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
