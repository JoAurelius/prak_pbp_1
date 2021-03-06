package controllers

import (
	"encoding/json"
	"net/http"
)

func SendErrorResponse(status int, errMsg string, header int, w http.ResponseWriter) {
	var response GeneralResponse
	response.Status = status
	response.Message = errMsg
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(header)
	json.NewEncoder(w).Encode(response)
}

func SendSuccessResponse(status int, msg string, header int, w http.ResponseWriter) {
	var response GeneralResponse
	response.Status = status
	response.Message = msg
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
