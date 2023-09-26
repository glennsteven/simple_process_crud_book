package controllers

import (
	"asis_quest/helper"
	"asis_quest/models"
	"asis_quest/presentation"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// Book is get all data book from database
func Book(w http.ResponseWriter, _ *http.Request) {
	var (
		book     []models.Books
		response presentation.ResponseList
	)

	if err := models.DB.Find(&book).Error; err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, fetchBook := range book {
		rupiah := helper.Accounting(fetchBook.Price)
		resp := presentation.Data{
			ID:          fetchBook.ID,
			Title:       fetchBook.Title,
			Description: fetchBook.Description,
			Category:    nil,
			Keyword:     fetchBook.Keyword,
			Price:       rupiah,
			Stock:       fetchBook.Stock,
			Publisher:   fetchBook.Publisher,
			CreatedAt:   fetchBook.CreatedAt,
			UpdatedAt:   fetchBook.UpdatedAt,
		}

		for _, bc := range fetchBook.BookCategory {
			var category models.Categories
			if err := models.DB.Where("id = ?", bc.CategoryID).Find(&category).Error; err != nil {
				helper.ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}

			resp.Category = append(resp.Category, category.Category)
		}
		response.Data = append(response.Data, resp)
	}

	response.Code = http.StatusOK
	response.Message = "Fetch all books"
	helper.ResponseJSON(w, http.StatusOK, response)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var (
		book    models.Books
		payload presentation.BookPayload
	)

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&payload); err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()
	var existingBook models.Books
	if err := models.DB.Where("title = ?", payload.Title).First(&existingBook).Error; err != nil {
		response := map[string]string{"message": "data book already exist"}
		helper.ResponseJSON(w, http.StatusConflict, response)
		return
	}

	book = models.Books{
		Title:       payload.Title,
		Description: payload.Description,
		Keyword:     payload.Keyword,
		Price:       payload.Price,
		Stock:       payload.Stock,
		Publisher:   payload.Publisher,
		CreatedAt:   time.Now().Format(`2006-01-02 15:04:05`),
		UpdatedAt:   time.Now().Format(`2006-01-02 15:04:05`),
	}

	if err := models.DB.Create(&book).Error; err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, categoryName := range payload.Categories {
		var category models.Categories
		if err := models.DB.Where("category = ?", categoryName).First(&category).Error; err != nil {
			category = models.Categories{
				Category:  categoryName,
				CreatedAt: time.Now().Format(`2006-01-02 15:04:05`),
				UpdatedAt: time.Now().Format(`2006-01-02 15:04:05`),
			}
			if err := models.DB.Create(&category).Error; err != nil {
				helper.ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
		bookCategory := models.BookCategory{
			BookID:     book.ID,
			CategoryID: category.ID,
		}

		if err := models.DB.Create(&bookCategory).Error; err != nil {
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	rupiah := helper.Accounting(book.Price)

	response := presentation.Response{
		Code:    http.StatusCreated,
		Message: "Successfully create new book",
		Data: presentation.Data{
			ID:          book.ID,
			Title:       book.Title,
			Description: book.Description,
			Keyword:     book.Keyword,
			Price:       rupiah,
			Stock:       book.Stock,
			Publisher:   book.Publisher,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
		},
	}

	response.Data.Category = append(response.Data.Category, payload.Categories...)

	helper.ResponseJSON(w, http.StatusCreated, response)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var book models.Books
	bookID := mux.Vars(r)["book_id"]
	if err := models.DB.Where("id = ?", bookID).First(&book).Error; err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.ResponseJSON(w, http.StatusOK, book)
}
