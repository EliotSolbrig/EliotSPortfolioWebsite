package router

import (
    "fmt"
    "net/http"
    "html/template"
    // "encoding/json"
    // "io"
    "strconv"
)

type ProjectPageRequest struct {
    TypeID int `json"type_id"`
}

func (router *Router) ProjectsPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("/projects")

    // var typeID int = 0

    // resBody,err := io.ReadAll(r.Body)
    // if err != nil {
    //     fmt.Println("Error reading request body in r_projects.go: ", err)
    // }
    // fmt.Println("resBody: ", string(resBody))
    //
    var projectRequest ProjectPageRequest
    // err = json.Unmarshal(resBody, &projectRequest)
    //
    // if err != nil {
    //     fmt.Println("Error unmarshaling projects page in r_projects.go: ", err)
    //     projectRequest.TypeID = 0
    // }
    projectRequest.TypeID,_ = strconv.Atoi(r.FormValue("project-type"))

    fmt.Println("projectRequest: ", projectRequest)

    projects,err := router.Service.GetAllProjects(r.Context(), projectRequest.TypeID)

    if err != nil {
        panic(fmt.Errorf("Error getting projects in r_projects.go: %s", err))
    }
    fmt.Println("projects: ", projects)


    templates := append([]string{"templates/Pages/Projects.html","templates/Components/ProjectPreview.html","templates/Pages/ProjectsPage.html",},basePasefiles...)
    tmpl := template.Must(template.ParseFiles(templates...))
    err = tmpl.ExecuteTemplate(w, "base", map[string]any{
        "Projects": projects,
    })
    if err != nil {
        panic(fmt.Errorf("Error executing Projects.html template in r_projects.go: %s", err))
    }
}

func (router *Router) GetProjectsPage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("/projectspage")

    // var typeID int = 0

    // resBody,err := io.ReadAll(r.Body)
    // if err != nil {
    //     fmt.Println("Error reading request body in r_projects.go: ", err)
    // }
    // fmt.Println("resBody: ", string(resBody))
    //
    var projectRequest ProjectPageRequest
    // err = json.Unmarshal(resBody, &projectRequest)
    //
    // if err != nil {
    //     fmt.Println("Error unmarshaling projects page in r_projects.go: ", err)
    //     projectRequest.TypeID = 0
    // }
    projectRequest.TypeID,_ = strconv.Atoi(r.FormValue("project-type"))

    fmt.Println("projectRequest: ", projectRequest)

    projects,err := router.Service.GetAllProjects(r.Context(), projectRequest.TypeID)

    if err != nil {
        panic(fmt.Errorf("Error getting projects in r_projects.go: %s", err))
    }
    fmt.Println("projects: ", projects)


    // templates := append([]string{"templates/Pages/Projects.html","templates/Components/ProjectPreview.html","templates/Pages/ProjectsPage.html",},basePasefiles...)
    templates := []string{"templates/Pages/ProjectsPage.html","templates/Components/ProjectPreview.html",}
    tmpl := template.Must(template.ParseFiles(templates...))
    err = tmpl.ExecuteTemplate(w, "projectpage", map[string]any{
        "Projects": projects,
    })
    if err != nil {
        panic(fmt.Errorf("Error executing Projects.html template in r_projects.go: %s", err))
    }

}
