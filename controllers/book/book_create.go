package controllers

import (
	"asis_quest/config/consts"
	"asis_quest/helper"
	"asis_quest/models"
	"asis_quest/presentation"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func (n *newBook) StoreBook(w http.ResponseWriter, r *http.Request) {
	var payload presentation.BookPayload

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&payload); err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	// Check if the book already exists
	existingBook, err := n.book.FindOne(models.Books{Title: payload.Title})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if existingBook != nil {
		response := map[string]string{"message": "Book already exists"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	book := models.Books{
		Title:       payload.Title,
		Description: payload.Description,
		Keyword:     payload.Keyword,
		Price:       payload.Price,
		Stock:       payload.Stock,
		Publisher:   payload.Publisher,
		CreatedAt:   time.Now().Format(consts.FormatDateIDN),
		UpdatedAt:   time.Now().Format(consts.FormatDateIDN),
	}

	// Store the new book
	newBook, err := n.book.Store(book)
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, categoryName := range payload.Categories {
		// Check if the category already exists
		category := models.Categories{Category: categoryName}
		existingCategory, err := n.ctg.FindOne(category)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if existingCategory == nil {
			// Category doesn't exist, create it
			categories, err := n.ctg.Store(category)
			if err != nil {
				helper.ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			category = *categories
		} else {
			// Category already exists, use the existing one
			category = *existingCategory
		}

		bookCategory := models.BookCategory{
			BookID:     newBook.ID,
			CategoryID: category.ID,
		}

		_, err = n.ctg.StoreBookCategory(bookCategory)
		if err != nil {
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	rupiah := helper.Accounting(newBook.Price)

	response := presentation.Response{
		Code:    http.StatusCreated,
		Message: "Successfully created a new book",
		Data: presentation.Data{
			ID:          newBook.ID,
			Title:       newBook.Title,
			Description: newBook.Description,
			Keyword:     newBook.Keyword,
			Price:       rupiah,
			Stock:       newBook.Stock,
			Publisher:   newBook.Publisher,
			CreatedAt:   newBook.CreatedAt,
			UpdatedAt:   newBook.UpdatedAt,
			Category:    payload.Categories,
		},
	}

	helper.ResponseJSON(w, http.StatusCreated, response)
}
