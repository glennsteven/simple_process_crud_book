package controllers

import (
	"asis_quest/repositories"
	"net/http"
)

type LibraryBook interface {
	StoreBook(w http.ResponseWriter, r *http.Request)
	DetailBook(w http.ResponseWriter, r *http.Request)
	ListBook(w http.ResponseWriter, _ *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
}

type newBook struct {
	book repositories.Book
	ctg  repositories.Category
}

func NewBook(
	book repositories.Book,
	ctg repositories.Category,
) LibraryBook {
	return &newBook{
		book: book,
		ctg:  ctg,
	}
}
