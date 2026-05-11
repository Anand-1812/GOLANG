package handlers

import (
	"fmt"
	"net/http"

	"http-handler/models"
)

func List(db models.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	}
}

func Price(db models.Database) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		item := r.URL.Query().Get("item")

		price, ok := db[item]

		if !ok {

			w.WriteHeader(http.StatusNotFound)

			fmt.Fprintf(w, "no such item: %q\n", item)

			return
		}

		fmt.Fprintf(w, "%s\n", price)
	}
}

