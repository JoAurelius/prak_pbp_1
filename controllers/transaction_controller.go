package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT * from transaction"

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
		var response TransactionResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = transaction
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		var response GeneralResponse
		response.Status = 204
		response.Message = "Error array is empty"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(response)
	}

}
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	transactionId := vars["transaction_id"]

	_, errQuery := db.Exec("DELETE FROM transaction WHERE id=?", transactionId)
	if errQuery == nil {
		sendSuccessResponse(w)
	} else {
		EmptyArrayErrorResponse(w)
	}
}
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		QeuryErrorResponse(w)
		return
	}
	vars := mux.Vars(r)
	var transaction Transaction = GetTransaction(vars["transaction_id"], w)
	var userid, _ = strconv.Atoi(r.Form.Get("age"))
	var productid, _ = strconv.Atoi(r.Form.Get("age"))
	var qyt, _ = strconv.Atoi(r.Form.Get("age"))
	if userid != 0 {
		transaction.UserID = userid
	}
	if productid != 0 {
		transaction.ProductID = productid
	}
	if qyt != 0 {
		transaction.Quantity = qyt
	}

	result, _ := db.Exec("UPDATE transaction SET userID = ?, productId = ?, quantity = ? WHERE id = ?", transaction.ProductID, transaction.ProductID, transaction.Quantity, transaction.ID)

	num, _ := result.RowsAffected()

	var response TransactionResponse
	if num != 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data = transaction
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		QeuryErrorResponse(w)
	}
}
func InsertTransaction(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		return
	}
	var transaction Transaction
	transaction.UserID, _ = strconv.Atoi(r.Form.Get("user_id"))
	transaction.ProductID, _ = strconv.Atoi(r.Form.Get("product_id"))
	transaction.Quantity, _ = strconv.Atoi(r.Form.Get("quantity"))

	result, _ := db.Exec("insert into transaction (UserID, ProductID, Quantity) values (?, ?, ?)",
		transaction.UserID, transaction.ProductID, transaction.Quantity)

	num, _ := result.RowsAffected()

	var response TransactionResponse
	if num != 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data = transaction
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		QeuryErrorResponse(w)
	}
}
func GetTransaction(transaction_id string, w http.ResponseWriter) Transaction {
	db := Connect()
	defer db.Close()
	var transaction Transaction
	query := "SELECT * from transaction WHERE ID = " + transaction_id
	rows, err := db.Query(query)
	if err != nil {
		QeuryErrorResponse(w)
	}
	for rows.Next() {
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Quantity); err != nil {
			QeuryErrorResponse(w)
		}
	}
	return transaction
}
