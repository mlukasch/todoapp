package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleAsNotFound(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(err.Error()))
	log.Printf("Error : %s", err.Error())
}

func HandleJsonResponse(response map[string]interface{}, w http.ResponseWriter) {
	data, err := json.Marshal(response)
	if err != nil {
		HandleAsNotFound(err, w)
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
}
