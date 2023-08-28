package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"

	"github.com/ahsanaminofficial/golang-project-careem/pkg/controllers"
)

func main() {
	fmt.Println("Starting the program")

	router := httprouter.New()
	router.GET("/", controllers.HomepageHandler)
	router.GET("/reviews", controllers.AllReviews)
	router.GET("/reviews/:id", controllers.ReviewById)
	router.POST("/review", controllers.CreateReview)

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		fmt.Println("Count not start server")
	}
}
