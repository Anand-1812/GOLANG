package main

import (
	"log"
	"net/http"

	"http-handler/models"
	"http-handler/handlers"
)

func main() {

	db := models.Database{
		"shoes": 50,
		"socks": 5,
		"hat":   20,
	}

	http.HandleFunc("/list", handlers.List(db))
	http.HandleFunc("/price", handlers.Price(db))

	log.Println("server running on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
