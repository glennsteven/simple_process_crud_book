package controllers

import (
	"asis_quest/config/consts"
	"asis_quest/helper"
	"asis_quest/models"
	"asis_quest/presentation"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func (n *newBook) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var payload presentation.BookPayload
	bookID := mux.Vars(r)["id"]
	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		return
	}
	findBook, err := n.book.FindID(models.Books{ID: int64(idBook)})
	if err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if findBook == nil {
		response := map[string]string{"message": "Data not found"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
		return
	}

	changeData := models.Books{
		Title:       payload.Title,
		Description: payload.Description,
		Keyword:     payload.Keyword,
		Price:       payload.Price,
		Stock:       payload.Stock,
		Publisher:   payload.Publisher,
		UpdatedAt:   time.Now().Format(consts.FormatDateIDN),
	}

	update, err := n.book.Update(findBook.ID, changeData)
	if err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.ResponseJSON(w, http.StatusOK, update)

}
