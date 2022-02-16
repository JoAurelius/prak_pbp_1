package controllers

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Println(err)
	}

	var user User
	var users []User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address); err != nil {
			log.Fatal(err.Error())
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

	} else {
		var response ErrorResponse
		response.Status = 204
		response.Message = "Error array is empty"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
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
	user.Name = r.Form.Get("namme")
	user.Age, _ = strconv.Atoi(r.Form.Get("age"))
	user.Address = r.Form.Get("address")

	_, errQuery := db.Exec("INSER INTO users(name,age,address) values (?,?,?)", user.Name, user.Age, user.Address)

	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 400
		response.Message = "Insert Failed!"
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
	fmt.Print(errQuery)
	if errQuery != nil {
		// sendSuccessResponse(w)
		fmt.Println()
	} else {
		sendErrorResponse(w)
	}
}
