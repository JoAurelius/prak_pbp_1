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
		QeuryErrorResponse(w)
		return
	}

	var user User
	var users []User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address); err != nil {
			QeuryErrorResponse(w)
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
		QeuryErrorResponse(w)
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
		QeuryErrorResponse(w)
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
		sendSuccessResponse(w)
	} else {
		QeuryErrorResponse(w)
	}
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		QeuryErrorResponse(w)
		return
	}
	var user User
	user.Name = r.Form.Get("name")
	user.Age, _ = strconv.Atoi(r.Form.Get("age"))
	user.Address = r.Form.Get("address")
	vars := mux.Vars(r)
	userID := vars["user_id"]

	if user.Name == "" {
		user.Name = GetUser(userID, w).Name
	}
	if user.Age == 0 {
		user.Age = GetUser(userID, w).Age
	}
	if user.Address == "" {
		user.Address = GetUser(userID, w).Address
	}

	result, _ := db.Exec("UPDATE users SET name = ?, age = ?, address = ? WHERE id = ?", user.Name, user.Age, user.Address, userID)

	num, _ := result.RowsAffected()

	var response UserResponse
	if num != 0 {
		response.Status = 200
		response.Message = "Success"
		response.Data = user
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		QeuryErrorResponse(w)
	}
}
func GetUser(user_id string, w http.ResponseWriter) User {
	db := Connect()
	defer db.Close()
	var user User
	query := "SELECT * from users WHERE ID = " + user_id
	rows, err := db.Query(query)
	if err != nil {
		QeuryErrorResponse(w)
	}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address); err != nil {
			QeuryErrorResponse(w)
		}
	}
	return user
}
