package repositories

import "asis_quest/models"

type User interface {
	FindUser(where models.User) (*models.User, error)
	InsertUser(insert models.User) error
}

type Book interface {
	Find(where []models.Books) ([]models.Books, error)
	Store(insert models.Books) (*models.Books, error)
	FindOne(where models.Books) (*models.Books, error)
	FindID(where models.Books) (*models.Books, error)
}

type Category interface {
	Store(insert models.Categories) (*models.Categories, error)
	StoreBookCategory(insert models.BookCategory) (*models.BookCategory, error)
	FindOne(where models.Categories) (*models.Categories, error)
	FindBookCategory(where models.BookCategory) ([]models.BookCategory, error)
	FindID(where models.Categories) (*models.Categories, error)
}
