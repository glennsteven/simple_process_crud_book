package httpx

import (
	"asis_quest/httpx/router"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RUN() {
	r := mux.NewRouter()
	router.Router(r)

	log.Fatal(http.ListenAndServe(":8000", r))
}
