package router

import (
    "fmt"
    "net/http"
    "html/template"
)

func (router *Router) SkillsPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("/skills")

    templates := append([]string{"templates/Pages/Skills.html",},basePasefiles...)
    tmpl := template.Must(template.ParseFiles(templates...))
    err := tmpl.ExecuteTemplate(w, "base", map[string]any{})
    if err != nil {
        panic(fmt.Errorf("Error executing Skills.html template in r_skills.go: %s",err))
    }
}
