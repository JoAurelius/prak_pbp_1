package controllers

import (
	"encoding/json"
	"net/http"
)

func EmptyArrayErrorResponse(w http.ResponseWriter) {
	var response GeneralResponse
	response.Status = 204
	response.Message = "Error array is empty"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(w http.ResponseWriter) {
	var response GeneralResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func QeuryErrorResponse(w http.ResponseWriter) {
	var response GeneralResponse
	response.Status = 400
	response.Message = "Error array is empty"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}
