package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT * from products"

	rows, err := db.Query(query)
	if err != nil {
		QeuryErrorResponse(w)
		return
	}

	var product Product
	var products []Product
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			QeuryErrorResponse(w)
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
		EmptyArrayErrorResponse(w)
	}

}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		QeuryErrorResponse(w)
		return
	}

	vars := mux.Vars(r)
	productID := vars["product_id"]

	_, errQuery := db.Exec("DELETE FROM products WHERE id=?", productID)
	if errQuery == nil {
		sendSuccessResponse(w)
	} else {
		EmptyArrayErrorResponse(w)
	}
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		QeuryErrorResponse(w)
		return
	}
	var product Product
	product.Name = r.Form.Get("name")
	product.Price, _ = strconv.Atoi(r.Form.Get("age"))
	vars := mux.Vars(r)
	productID := vars["product_id"]

	if product.Name == "" {
		product.Name = GetProduct(productID, w).Name
	}
	if product.Price == 0 {
		product.Price = GetProduct(productID, w).Price
	}

	result, _ := db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, productID)

	num, _ := result.RowsAffected()

	var response ProductResponse
	if num != 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data = product
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		QeuryErrorResponse(w)
	}
}
func InserProduct(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		return
	}
	var product Product
	product.Name = r.Form.Get("name")
	product.Price, _ = strconv.Atoi(r.Form.Get("price"))

	result, _ := db.Exec("insert into products (name, price) values (?, ?)", product.Name, product.Price)

	num, _ := result.RowsAffected()

	var response ProductResponse
	if num != 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data = product
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		QeuryErrorResponse(w)
	}
}
func GetProduct(product_id string, w http.ResponseWriter) Product {
	db := Connect()
	defer db.Close()
	var product Product
	query := "SELECT * from products WHERE ID = " + product_id
	rows, err := db.Query(query)
	if err != nil {
		QeuryErrorResponse(w)
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			QeuryErrorResponse(w)
		}
	}
	return product
}
