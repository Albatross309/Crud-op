package main

import (
	"crud-app/database"
	"crud-app/routes"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Database
	database.ConnectDB()

	e := echo.New()

	// Define Routes
	routes.UserRoutes(e)

	fmt.Println("Server started at :1323")
	log.Fatal(e.Start(":1323"))
}

// package main

// import (
// 	userServices "crud-app/services"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	// services "command-line-argumentsD:\\New folder\\Golang\\Crud-op\\services\\userServices.go"
// 	_ "github.com/go-sql-driver/mysql"
// 	// "github.com/labstack/echo/middleware"
// 	"github.com/labstack/echo/v4"
// )

// func main(e *echo.Echo) {
// 	// e.Use(middleware.Logger())
// 	// e.Use(middleware.Recover())
// 	e.GET("/users", userServices.getUser)
// 	e.GET("/users", userServices.getAllUsers)
// 	e.POST("/users", userServices.createUser)
// 	e.PUT("/users", userServices.updateUser)
// 	e.DELETE("/users", userServices.deleteUser)
// 	// var (
// 	// 	users = map[int]*user{}
// 	// 	seq   = 1
// 	// 	lock  = sync.Mutex{}
// 	// )

// 	// router := mux.NewRouter()
// 	// router.HandleFunc("/users", userServices.createUser).Methods(http.MethodPost)
// 	// router.HandleFunc("/users", userServices.getAllUsers).Methods(http.MethodGet)
// 	// router.HandleFunc("/users", userServices.getAllUsers).Methods(http.MethodGet)
// 	// router.HandleFunc("/users", userServices.updateUser).Methods(http.MethodPut)
// 	// router.HandleFunc("/users", userServices.deleteUser).Methods(http.MethodDelete)

// 	fmt.Print("Starting server at port")
// 	log.Fatal(http.ListenAndServe(":1323", e))
// 	// e.Logger.Fatal(e.Start(":1323"))
// }
