package main

import (
	"asis_quest/authcontroller"
	"asis_quest/controllers"
	"asis_quest/middlewares"
	"asis_quest/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()

	// GET All data book
	api.HandleFunc("/books", controllers.Book).Methods("GET")

	// GET data book by ID
	api.HandleFunc("/book/{book_id}", controllers.GetBook).Methods("GET")

	// POST data book for create new book
	api.HandleFunc("/book", controllers.CreateBook).Methods("POST")

	api.Use(middlewares.JWTMiddleware)
	log.Fatal(http.ListenAndServe(":8000", r))
}
