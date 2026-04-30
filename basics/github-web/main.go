package main

import (
	"log"
	"net/http"
	"os"

	"github-web/github"
	"github-web/web"
)

func main() {
	err := github.LoadData("repo:golang/go is:issue bug")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", web.IndexHandler)
	http.HandleFunc("/issue", web.IssueHandler)
	http.HandleFunc("/users", web.UsersHandler)
	http.HandleFunc("/milestones", web.MilestonesHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
