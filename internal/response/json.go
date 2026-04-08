package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, map[string]string{
		"error": message,
	})
}
