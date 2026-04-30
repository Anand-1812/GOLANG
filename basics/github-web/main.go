package main

import (
	"log"
	"net/http"

	"github-web/github"
	"github-web/web"
)

func main() {
	err := github.LoadData("repo:golang/go+is:issue+bug")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", web.IndexHandler)
	http.HandleFunc("/issue", web.IssueHandler)
	http.HandleFunc("/users", web.UsersHandler)
	http.HandleFunc("/milestones", web.MilestonesHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
