package controllers

import (
	"encoding/json"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter) {
	var response ErrorResponse
	response.Status = 204
	response.Message = "Error array is empty"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(user User, w http.ResponseWriter) {
	var response UserResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
