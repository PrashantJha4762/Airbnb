package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter,status int , data any) error{
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}