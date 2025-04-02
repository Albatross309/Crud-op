package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// func DbConnection() (*sql.DB, error) {
// 	connectionString := "root:albatross@/devbook?charset=utf8&parseTime=True&loc=Local"

// 	db, err := sql.Open("mysql", connectionString)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:albatross@/devbook?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Failed to connect to the database!", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Database is not reachable!", err)
	}
	fmt.Println("Connected to MySQL database!")
}
