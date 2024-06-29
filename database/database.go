package database

import (
    "database/sql"
    "fmt"
    "context"

    _ "github.com/mattn/go-sqlite3"
)

type IDatabase interface{
    GetAllJobs(ctx context.Context) ([]Job, error)
    GetExperienceItemsByID(ctx context.Context, jobID int) ([]ExperienceItem, error)
    GetAllSchools(ctx context.Context) ([]*School, error)
    GetSchoolClassesFromSchoolID(ctx context.Context, schoolID int) ([]Class, error)
    GetAllProjects(ctx context.Context) ([]*Project, error)
    GetProjectImagesFromID(ctx context.Context, projectID int) ([]*ProjectImage, error)
}

type Database struct {
	db *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

func GetDB() *sql.DB {
    fmt.Println("GetDB()")
    // dbFile,err := os.Create("database.db")
    // if err != nil {
    //     panic(fmt.Errorf("Error creating sqlite database file: %s", err))
    // }
    // fmt.Println("dbFile: ", dbFile)
    db,err := sql.Open("sqlite3", "./database.db")
        if err != nil {
            panic(fmt.Errorf("Error opening sqlite database in database.go: %s", err))
    }
    return db
}

