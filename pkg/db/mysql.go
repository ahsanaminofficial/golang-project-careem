package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() *sql.DB {
	database, err := sql.Open("mysql", "root:ahsan1307@tcp(127.0.0.1:3306)/reviews_db")
	if err != nil {
		fmt.Println("Cannot connect to mysql Database")
	}

	fmt.Println("Database Connected")
	// defer db.Close()
	db = database
	return database
}

func GetDB() *sql.DB {
	return db
}
