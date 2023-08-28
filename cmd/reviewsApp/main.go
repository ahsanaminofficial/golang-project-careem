package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"

	"github.com/ahsanaminofficial/golang-project-careem/tree/main/pkg/controllers"
)

func main() {
	fmt.Println("Starting the program")

	router := httprouter.New()
	router.GET("/", controllers.HomepageHandler)
}
