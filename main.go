package main

import (
    "fmt"
    "os"
    "net/http"
	"github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "main/router"
    "main/service"
    "main/database"
    "main/Util"
)

const staticDir string = "/static/"

func main(){
    fmt.Println("main.go")
    fmt.Println("\n-------------------------------------------------------\n")

    godotenv.Load()

    hostname := Util.GetHostname()
    fmt.Println("hostname: ", hostname)
    port := os.Getenv("PORT")
    fmt.Println("port: ", port)

    r := mux.NewRouter()

    router := router.NewRouter(service.NewService(database.NewDatabase(database.GetDB())))

    staticHandler := http.StripPrefix(staticDir, http.FileServer(http.Dir("static/")))
    r.PathPrefix(staticDir).Handler(staticHandler)

    r.HandleFunc("/", router.IndexPage)
    r.HandleFunc("/workexperience", router.WorkExperiencePage)
    r.HandleFunc("/education", router.EducationPage)
    r.HandleFunc("/projects", router.ProjectsPage)
    r.HandleFunc("/projectpage", router.GetProjectsPage)
    r.HandleFunc("/interests", router.InterestsPage)

    fmt.Println("\nListening on ", port)
    panic(http.ListenAndServe(":" + port, r))


}
