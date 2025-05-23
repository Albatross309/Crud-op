package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	// connStr := "host=localhost port=1323 user=postgres password=postgres dbname=testdb sslmode=disable"
	DB, err = sql.Open("mysql", "root:albatross@/devbook?charset=utf8&parseTime=True&loc=Local")

	// sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database! ", err)
	}
	// if err = DB.Ping(); err != nil {
	// 	log.Fatal("Database is not reachable! ", err)
	// }
	fmt.Println("Connected to MySQL database! ")
}
