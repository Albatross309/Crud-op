package main

import (
	userServices "crud-app/services"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", userServices.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", userServices.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", userServices.GetUserByID).Methods(http.MethodGet)
	router.HandleFunc("/users", userServices.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users", userServices.DeleteUser).Methods(http.MethodDelete)

	fmt.Print("Listenning on 1323")
	log.Fatal(http.ListenAndServe(":1323", router))
}
