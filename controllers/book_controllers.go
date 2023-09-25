package controllers

import (
	"asis_quest/helper"
	"asis_quest/models"
	"asis_quest/repositories"
	"encoding/json"
	"net/http"
	"time"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Categories

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&category)
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	category.DeletedAt = nil

	if err := models.DB.Create(&category).Error; err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.ResponseJSON(w, http.StatusCreated, category)
}

func Book(w http.ResponseWriter, _ *http.Request) {
	var book []models.Books

	if err := models.DB.Find(&book).Error; err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.ResponseJSON(w, http.StatusOK, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var (
		book     models.Books
		response repositories.Response
	)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	var exists bool

	if err := models.DB.Where("title = ?", book.Title).Find(&book).Error; err != nil {
		if !exists {
			response := map[string]string{"message": "data book already exist"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}
	} else {
		if err := models.DB.Create(&book).Error; err != nil {
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	rupiah := helper.Accounting(book.Price)

	response.Code = http.StatusCreated
	response.Message = "Successfully create new book"
	response.Data.ID = book.Id
	response.Data.CategoryID = book.CategoryID
	response.Data.Title = book.Title
	response.Data.Category = book.Category
	response.Data.Keyword = book.Keyword
	response.Data.Price = rupiah
	response.Data.Stock = book.Stock
	response.Data.Publisher = book.Publisher
	response.Data.Description = book.Description
	response.Data.CreatedAt = book.CreatedAt
	response.Data.UpdatedAt = book.UpdatedAt
	response.Data.DeletedAt = book.DeletedAt

	helper.ResponseJSON(w, http.StatusCreated, response)
}
