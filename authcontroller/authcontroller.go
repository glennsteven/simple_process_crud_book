package authcontroller

import (
	"asis_quest/config"
	"asis_quest/helper"
	"asis_quest/models"
	"asis_quest/presentation"
	"asis_quest/repositories"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"time"
)

type newLogin struct {
	user repositories.User
}

func NewUser(user repositories.User) User {
	return &newLogin{user: user}
}

func (n *newLogin) Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	findUser, err := n.user.FindUser(userInput)
	if err != nil {
		switch err.Error() {
		case "not found":
			response := map[string]string{"message": "invalid username or password"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	err = compareHashPassword([]byte(findUser.Password), []byte(userInput.Password))
	if err != nil {
		response := map[string]string{"message": "invalid username or password"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	exp := time.Now().Add(time.Hour * 5)
	claims := &config.JWTclaim{
		Username: findUser.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "test_asia_quest",
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	tokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgorithm.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"token": token}
	helper.ResponseJSON(w, http.StatusOK, response)
	return
}

func (n *newLogin) Register(w http.ResponseWriter, r *http.Request) {
	var (
		user models.User
	)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(r.Body)

	hashPassword, err := generateHashPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	user.Password = string(hashPassword)

	err = n.user.InsertUser(user)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := presentation.ResponseRegister{
		Code:    http.StatusCreated,
		Message: "registered successfully",
		Data: presentation.DataRegister{
			FullName: user.FullName,
			Username: user.UserName,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, response)
}
