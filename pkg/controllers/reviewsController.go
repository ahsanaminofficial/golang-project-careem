package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HomepageHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	io.WriteString(w, "Home page")
	fmt.Println("homepage handler")
}

func AllReviews(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	io.WriteString(w, "AllReviews")
	fmt.Println("AllReviews")
}

func ReviewById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	io.WriteString(w, "ReviewById")
	fmt.Println("ReviewById")
}

func CreateReview(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	io.WriteString(w, "CreateReview")
	fmt.Println("CreateReview")
}
