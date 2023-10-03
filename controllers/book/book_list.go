package controllers

import (
	"asis_quest/helper"
	"asis_quest/models"
	"asis_quest/presentation"
	"net/http"
)

func (n *newBook) ListBook(w http.ResponseWriter, _ *http.Request) {
	var (
		book     []models.Books
		response presentation.ResponseList
	)

	findBooks, err := n.book.Find(book)
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, fetchBook := range findBooks {
		rupiah := helper.Accounting(fetchBook.Price)
		resp := presentation.Data{
			ID:          fetchBook.ID,
			Title:       fetchBook.Title,
			Description: fetchBook.Description,
			Keyword:     fetchBook.Keyword,
			Price:       rupiah,
			Stock:       fetchBook.Stock,
			Publisher:   fetchBook.Publisher,
			CreatedAt:   fetchBook.CreatedAt,
			UpdatedAt:   fetchBook.UpdatedAt,
		}

		whereIDBC := models.BookCategory{BookID: fetchBook.ID}
		bookCategory, err := n.ctg.FindBookCategory(whereIDBC)
		if err != nil {
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}

		for _, c := range bookCategory {
			category, err := n.ctg.FindID(models.Categories{ID: c.CategoryID})
			if err != nil {
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
