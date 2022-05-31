package main

import (
	"log"
	"net/http"
	"packages/auth"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	// Init the mux router
	router := mux.NewRouter()

	router.HandleFunc("/login", auth.Login).Methods("GET")
	router.HandleFunc("/register", Register).Methods("POST")
	router.HandleFunc("/products", GetAllProd).Methods("GET")
	router.HandleFunc("/movies/{movieid}", DeleteMovie).Methods("DELETE")
	router.HandleFunc("/movies/", DeleteMovies).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
