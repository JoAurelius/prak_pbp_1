package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT * from products"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var product Product
	var products []Product
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			log.Fatal(err.Error())
		} else {
			products = append(products, product)
		}
	}
	if len(products) > 1 {
		var response ProductsResponse
		response.Status = 200
		response.Message = "Sucess"
		response.Data = products
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else if len(products) == 1 {
		var response ProductResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = product
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		var response ErrorResponse
		response.Status = 204
		response.Message = "Error array is empty"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	}

}
