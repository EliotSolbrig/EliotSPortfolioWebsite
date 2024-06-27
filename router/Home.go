package router

import (
    "fmt"
    "html/template"
    "net/http"
)

func (router *Router) IndexPage(w http.ResponseWriter, r *http.Request){
    templates := append([]string{"templates/index.html"}, basePasefiles...)

    tmpl := template.Must(template.ParseFiles(templates...))
    err := tmpl.ExecuteTemplate(w, "base", map[string]any{
        "Data": []any{},
    })
    if err != nil {
        panic(fmt.Errorf("Error executing base IndexPage base template: %s", err))
    }
}
