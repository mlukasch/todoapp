package utils

import (
	"log"
	"net/http"
)

func HandleAsNotFound(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(err.Error()))
	log.Printf("Error : %s", err.Error())
}
