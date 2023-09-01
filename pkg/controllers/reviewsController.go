package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ahsanaminofficial/golang-project-careem/pkg/models"
	"github.com/julienschmidt/httprouter"
)

func HomepageHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	res, _ := json.Marshal("Homepage")
	w.Write(res)
	fmt.Println("homepage handler")
}

func AllReviews(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("AllReviews handler")
	rows := models.GetAllReviews()

	var result []models.Review
	for rows.Next() {
		var row models.Review
		rows.Scan(&row.Id, &row.Body, &row.InstrutorName, &row.Rating)
		result = append(result, row)
	}

	res, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func ReviewById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("ReviewById")

	ReviewId := params.ByName("id")

	rows := models.GetReviewById(ReviewId)

	var result []models.Review
	for rows.Next() {
		var row models.Review
		rows.Scan(&row.Id, &row.Body, &row.InstrutorName, &row.Rating)
		result = append(result, row)
	}

	res, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateReview(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("Create Review")
	rating, _ := strconv.Atoi(r.FormValue("rating"))

	models.CreateReview(&models.Review{Id: r.FormValue("id"), Body: r.FormValue("body"), InstrutorName: r.FormValue("instructor_name"), Rating: rating})
}

func DeleteReviewByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("delete review")
	models.DeleteReviewByID(params.ByName(("id")))
}
