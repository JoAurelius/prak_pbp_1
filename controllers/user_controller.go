package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connect()
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
		var response UserResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = user
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
