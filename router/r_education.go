package router

import (
    "fmt"
    "net/http"
    "html/template"
)

func (router *Router) EducationPage(w http.ResponseWriter, r *http.Request){
    fmt.Println("/educationpage")

    schoolData,err := router.Service.GetAllSchools(r.Context())
    if err != nil {
        fmt.Println(fmt.Errorf("Error getting school data in r_education.go: ", err))
    }
    fmt.Println("schoolData: ", schoolData)

    templates := append([]string{"templates/Pages/Education.html","templates/Components/School.html","templates/Components/Class.html",},basePasefiles...)
    tmpl := template.Must(template.ParseFiles(templates...))
    err = tmpl.ExecuteTemplate(w, "base", map[string]any{
        "Data": map[string]any{
            "Schools": schoolData,
        },
    })
    if err != nil {
        panic(fmt.Errorf("Error executing Education.html template in r_education.go: %s", err))
    }
}

