package router

import (
	"asis_quest/authcontroller"
	"asis_quest/connection"
	controllers "asis_quest/controllers/book"
	"asis_quest/middlewares"
	"asis_quest/repositories"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Router(r *mux.Router) {
	var DB *gorm.DB
	database := connection.NewDatabase(DB)
	repoUser := repositories.NewUser(database)
	userLogin := authcontroller.NewUser(repoUser)

	// Path API & Method API
	r.HandleFunc("/login", userLogin.Login).Methods("POST")
	r.HandleFunc("/register", userLogin.Register).Methods("POST")

	// Sub-router
	api := r.PathPrefix("/library").Subrouter()

	repoBook := repositories.NewBook(database)
	repoCategory := repositories.NewCategory(database)
	book := controllers.NewBook(repoBook, repoCategory)

	// POST data book for create new book
	api.HandleFunc("/book",
		book.StoreBook,
	).Methods("POST")

	// GET data book by ID
	api.HandleFunc("/book/{book_id}",
		book.DetailBook,
	).Methods("GET")

	// GET All data book
	api.HandleFunc("/books",
		book.ListBook,
	).Methods("GET")

	// PUT data book by ID
	api.HandleFunc("/book/{id}",
		book.UpdateBook,
	).Methods("PUT")

	api.Use(middlewares.JWTMiddleware)
}
