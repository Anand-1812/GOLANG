package web

import (
	"html/template"
	"net/http"
	"strconv"

	"github-web/github"
)

var templates = template.Must(template.ParseGlob("web/template/*.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", github.CachedResult)
}

func IssueHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	for _, issue := range github.CachedResult.Items {
		if issue.Number == id {
			templates.ExecuteTemplate(w, "issue.html", issue)
			return
		}
	}

	http.NotFound(w, r)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	userMap := make(map[string]bool)
	var users []string

	for _, issue := range github.CachedResult.Items {
		if !userMap[issue.User.Login] {
			userMap[issue.User.Login] = true
			users = append(users, issue.User.Login)
		}
	}

	templates.ExecuteTemplate(w, "users.html", users)
}

func MilestonesHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "milestone.html", github.CachedResult.Items)
}
