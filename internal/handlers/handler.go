package handlers

import (
	"log"
	"net/http"
)

type Handler interface {
	Supports(url string) bool
	Handle(dir string, res http.ResponseWriter, req *http.Request)
}

func CreateSuccessResponse(res http.ResponseWriter) {
	// create response binary data
	data := []byte("success")

	// write `data` to response
	_, err := res.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
