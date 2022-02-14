package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * from products"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var transaction Transaction
	var transactions []Transaction
	for rows.Next() {
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Quantity); err != nil {
			log.Fatal(err.Error())
		} else {
			transactions = append(transactions, transaction)
		}
	}
	if len(transactions) > 1 {
		var response TransactionsResponse
		response.Status = 200
		response.Message = "Sucess"
		response.Data = transactions
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else if len(transactions) == 1 {
		var response Transactionesponse
		response.Status = 200
		response.Message = "Success"
		response.Data = transaction
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		var response ErrorResponse
		response.Status = 204
		response.Message = "Error array is empty"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}
