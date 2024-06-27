package router

import (
    "fmt"
    "net/http"
    "html/template"
)

type Job struct {
    Title string
    CompanyName string
    Location string
    StartDate string
    EndDate string
    Experience []string
}

type ExperienceItem struct {
    Text string
}

func (router *Router) WorkExperiencePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("/workexperience")

    takeda1 := Job {
        Title: "Production Engineer",
        CompanyName: "Takeda Pharmaceutical Company",
        Location: "Social Circle, Georgia, USA",
        Experience: []string{"Built/manage application in VBA which monitors process times in order to catch and correct process delays. Managed and facilitated user feedback to continuously improve the application","Reduced manufacturing events, user interventions, and batch record alerts","Troubleshooting of equipment, software, and process issues in GMP manufacturing facility for a batch biopharmaceutical process","Identified and developed opportunities to reduce cycle time, error proof operations, and eliminate waste.",},
    }

    jobs := []Job{takeda1,}

    templates := append([]string{"templates/Pages/WorkExperience.html","templates/Components/Job.html",},basePasefiles...)
    tmpl := template.Must(template.ParseFiles(templates...))
    err := tmpl.ExecuteTemplate(w, "base", map[string]any{
        "Experience": jobs,
    })
    if err != nil {
        panic(fmt.Errorf("Error executing WorkExperiences.html template in r_work_experience.go: %s", err))
    }
}
