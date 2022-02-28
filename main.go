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
	// router.HandleFunc("/transactions", controllers.GetAllTransaction).Methods("GET")
	router.HandleFunc("/transactions", controllers.GetAllDetailedTransaction).Methods("GET")
	router.HandleFunc("/transactions/{transaction_id}", controllers.GetDetailedTransaction).Methods("GET")
	router.HandleFunc("/users/{user_id}/transactions", controllers.GetDetailedTransactionFromUser).Methods("GET")
	router.HandleFunc("/transactions/users/{user_id}", controllers.GetDetailedTransactionFromUser).Methods("GET")

	router.HandleFunc("/users/{user_id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/products/{product_id}", controllers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/transactions/{transaction_id}", controllers.DeleteTransaction).Methods("DELETE")

	router.HandleFunc("/users", controllers.InsertNewUser).Methods("POST")
	router.HandleFunc("/products", controllers.InserProduct).Methods("POST")
	router.HandleFunc("/transactions", controllers.InsertTransaction).Methods("POST")
	router.HandleFunc("/users/{user_id}", controllers.Login).Methods("POST")

	router.HandleFunc("/users/{user_id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/products/{products_id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/transactions/{transactions_id}", controllers.UpdateTransaction).Methods("PUT")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
