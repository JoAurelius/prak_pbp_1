package main

import (
	"fmt"
	"log"
	"net/http"
	"prak_pbp_1/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/transactions", controllers.GetAllTransaction).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
