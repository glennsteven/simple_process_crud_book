package controllers

import (
	"asis_quest/helper"
	"asis_quest/models"
	"asis_quest/presentation"
	"errors"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (n *newBook) DetailBook(w http.ResponseWriter, r *http.Request) {
	var (
		response     presentation.Response
		bookCategory []models.BookCategory
	)
	bookID := mux.Vars(r)["book_id"]
	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		return
	}
	book := models.Books{ID: int64(idBook)}
	findBook, err := n.book.FindID(book)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if findBook == nil {
		response := map[string]string{"message": "Data not found"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	rupiah := helper.Accounting(findBook.Price)
	response.Code = http.StatusOK
	response.Message = "Successfully get book"
	response.Data = presentation.Data{
		ID:          findBook.ID,
		Title:       findBook.Title,
		Description: findBook.Description,
		Keyword:     findBook.Keyword,
		Price:       rupiah,
		Stock:       findBook.Stock,
		Publisher:   findBook.Publisher,
		CreatedAt:   findBook.CreatedAt,
		UpdatedAt:   findBook.UpdatedAt,
	}
	whereIDBC := models.BookCategory{BookID: findBook.ID}
	bookCategory, err = n.ctg.FindBookCategory(whereIDBC)
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, c := range bookCategory {
		whereIDC := models.Categories{ID: c.CategoryID}
		findCategory, err := n.ctg.FindID(whereIDC)
		if err != nil {
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
		response.Data.Category = append(response.Data.Category, findCategory.Category)
	}

	helper.ResponseJSON(w, http.StatusOK, response)
}
