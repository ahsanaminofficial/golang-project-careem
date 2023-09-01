package models

import (
	"database/sql"
	"fmt"

	mysql "github.com/ahsanaminofficial/golang-project-careem/pkg/db"
)

var db *sql.DB

type Review struct {
	Id            string `json:"id"`
	Body          string `json:"body"`
	InstrutorName string `json:"instructor_name"`
	Rating        int    `json:"rating"`
}

func init() {
	db = mysql.Connect()
	query := "CREATE TABLE IF NOT EXISTS reviews (`id` varchar(10), `body` varchar(100), `instructorName` varchar(40), `rating` int(2));"
	_, err := db.Query(query)

	if err != nil {
		fmt.Print(err)
	}
}

func GetAllReviews() *sql.Rows {
	query := "SELECT * FROM reviews"
	result, err := db.Query(query)

	if err != nil {
		fmt.Println("Some error while reading database")
	}

	fmt.Println(result)

	return result
}

func GetReviewById(id string) *sql.Rows {
	fmt.Println("id:", id)
	query := fmt.Sprintf("SELECT * FROM reviews WHERE id = %s;", id)
	result, err := db.Query(query)

	if err != nil {
		fmt.Println("Some error while reading database")
	}

	fmt.Println(result)

	return result
}

func CreateReview(r *Review) {
	query := fmt.Sprintf("INSERT INTO reviews (id, body, instructorName, rating) VALUES ('%s', '%s', '%s', '%d');", r.Id, r.Body, r.InstrutorName, r.Rating)
	result, err := db.Query(query)

	if err != nil {
		fmt.Println("Some error while creating the review:", err)
	}

	fmt.Println("Created:", result)
}

func DeleteReviewByID(id string) {
	query := fmt.Sprintf("DELETE FROM reviews WHERE id = '%s'", id)
	result, err := db.Query(query)

	if err != nil {
		fmt.Println("Some error while deleting the review:", err)
	}

	fmt.Println("Deleted:", result)
}
