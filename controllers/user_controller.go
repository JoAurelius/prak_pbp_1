package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT * from users"

	rows, err := db.Query(query)
	if err != nil {
		SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
		return
	}

	var user User
	var users []User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address); err != nil {
			SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
			return
		} else {
			users = append(users, user)
		}
	}
	if len(users) > 1 {
		var response UsersResponse
		response.Status = 200
		response.Message = "Sucess"
		response.Data = users
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else if len(users) == 1 {
		var response UserResponse
		response.Status = 200
		response.Message = "Sucess"
		response.Data = user
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
	}

}
func InsertNewUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		return
	}
	var user User
	user.Name = r.Form.Get("name")
	user.Age, _ = strconv.Atoi(r.Form.Get("age"))
	user.Address = r.Form.Get("address")

	result, _ := db.Exec("insert into users (name, age, address) values (?, ?, ?)", user.Name, user.Age, user.Address)

	num, _ := result.RowsAffected()

	var response UserResponse
	if num != 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data = user
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	userID := vars["user_id"]

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?", userID)

	if errQuery == nil {
		SendSuccessResponse(200, "Delete Success", http.StatusAccepted, w)
	} else {
		SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
	}
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
		return
	}

	vars := mux.Vars(r)
	var user User = GetUser(vars["user_id"], w)
	if (r.Form.Get("name")) != "" {
		user.Name = r.Form.Get("name")
	}
	if r.Form.Get("address") != "" {
		user.Address = r.Form.Get("address")
	}
	var temp, _ = strconv.Atoi(r.Form.Get("age"))
	if temp != 0 {
		user.Age, _ = strconv.Atoi(r.Form.Get("age"))
	}

	result, _ := db.Exec("UPDATE users SET name = ?, age = ?, address = ? WHERE id = ?", user.Name, user.Age, user.Address, user.ID)

	num, _ := result.RowsAffected()

	var response UserResponse
	if num != 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data = user
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
	}
}
func GetUser(user_id string, w http.ResponseWriter) User {
	db := Connect()
	defer db.Close()
	var user User
	query := "SELECT * from users WHERE ID = " + user_id
	rows, err := db.Query(query)
	if err != nil {
		SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
	}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address); err != nil {
			SendErrorResponse(404, "Query Error", http.StatusBadRequest, w)
		}
	}
	return user
}
