package router

import (
    "fmt"
    "net/http"
    "html/template"
)

func (router *Router) InterestsPage(w http.ResponseWriter,r *http.Request) {
    fmt.Println("/interests")

    templates := append([]string{"templates/Pages/Interests.html",},basePasefiles...)
    tmpl := template.Must(template.ParseFiles(templates...))
    err := tmpl.ExecuteTemplate(w, "base", map[string]any{})
    if err != nil {
        panic(fmt.Errorf("Error executing Interests.html template in r_interests.go: %s", err))
    }
}
