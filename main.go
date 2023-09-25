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

	api.HandleFunc("/category", controllers.CreateCategory).Methods("POST")
	api.HandleFunc("/book", controllers.Book).Methods("GET")
	api.HandleFunc("/book", controllers.CreateBook).Methods("POST")

	api.Use(middlewares.JWTMiddleware)
	log.Fatal(http.ListenAndServe(":8000", r))
}
