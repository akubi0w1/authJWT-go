package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func Success(w http.ResponseWriter, data interface{}) {
	jsonData, _ := json.Marshal(data)
	// w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	jsonData, err := json.Marshal(&Response{
		Code:    code,
		Message: message,
	})
	if err != nil {
		log.Fatal("json marshal error")
	}
	w.WriteHeader(code)
	// w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
