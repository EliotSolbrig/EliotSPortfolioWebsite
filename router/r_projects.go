package router

import (
    "fmt"
    "net/http"
    "html/template"
)

func (router *Router) ProjectsPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("/projects")

    projects,err := router.Service.GetAllProjects(r.Context())

    if err != nil {
        panic(fmt.Errorf("Error getting projects in r_projects.go: %s", err))
    }
    fmt.Println("projects: ", projects)


    templates := append([]string{"templates/Pages/Projects.html","templates/Components/ProjectPreview.html",},basePasefiles...)
    tmpl := template.Must(template.ParseFiles(templates...))
    err = tmpl.ExecuteTemplate(w, "base", map[string]any{
        "Projects": projects,
    })
    if err != nil {
        panic(fmt.Errorf("Error executing Projects.html template in r_projects.go: %s", err))
    }
}
