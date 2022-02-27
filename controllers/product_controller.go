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
		SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
		return
	}

	var product Product
	var products []Product
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
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
		SendErrorResponse(204, "Array Empty", http.StatusNoContent, w)
	}

}

// func DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	db := Connect()
// 	defer db.Close()

// 	err := r.ParseForm()
// 	if err != nil {
// 		QeuryErrorResponse(w)
// 		return
// 	}

// 	vars := mux.Vars(r)
// 	productID := vars["product_id"]

// 	_, errQuery := db.Exec("DELETE FROM products WHERE id=?", productID)
// 	if errQuery == nil {
// 		sendSuccessResponse(w)
// 	} else {
// 		EmptyArrayErrorResponse(w)
// 	}
// }
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
		return
	}
	vars := mux.Vars(r)
	var product Product = GetProduct(vars["product_id"], w)
	if (r.Form.Get("name")) != "" {
		product.Name = r.Form.Get("name")
	}
	var temp, _ = strconv.Atoi(r.Form.Get("price"))
	if temp != 0 {
		product.Price = temp
	}
	result, _ := db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, product.ID)

	num, _ := result.RowsAffected()

	var response ProductResponse
	if num != 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data = product
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
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
		SendErrorResponse(404, "Empty Array", http.StatusBadRequest, w)
	}
}
func GetProduct(product_id string, w http.ResponseWriter) Product {
	db := Connect()
	defer db.Close()
	var product Product
	query := "SELECT * from products WHERE ID = " + product_id
	rows, err := db.Query(query)
	if err != nil {
		SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
		}
	}
	return product
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
		return
	}

	vars := mux.Vars(r)
	productID := vars["product_id"]

	result, errQuery := db.Exec("DELETE FROM transaction WHERE ProductID = ?", productID)
	if errQuery != nil {
		SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
		return
	}
	num, _ := result.RowsAffected()

	if num != 0 {
		result, errQuery := db.Exec("DELETE FROM products WHERE id=?", productID)
		if errQuery == nil {
			SendSuccessResponse(200, "Delete Product Success", http.StatusAccepted, w)
		} else {
			SendErrorResponse(404, "Delete Query from Product Error", http.StatusBadRequest, w)
			return
		}
		num, _ := result.RowsAffected()
		if num != 0 {
			SendSuccessResponse(200, "Delete Product Success", http.StatusAccepted, w)
		} else {
			SendErrorResponse(204, "No Delete Row", http.StatusNoContent, w)
			return
		}
	} else {
		SendErrorResponse(204, "No Delete Row", http.StatusBadRequest, w)
	}
}
