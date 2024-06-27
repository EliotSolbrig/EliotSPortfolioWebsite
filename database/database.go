package database

import (
    "database/sql"
    "os"
    "fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type IDatabase interface {

}

type Database struct {
    DB *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
    return &Database{
        DB: db,
    }
}

func GetDB() *sql.DB {
	godotenv.Load()

    var user, pwd, dbHost, dbPort, dbName string

	// pwd := os.Getenv("Dconst B_PASSWORD")
    if os.Getenv("ENV_TYPE") == "dev" {
        pwd = ""
        user = os.Getenv("DEV_DB_USER")
        dbName = os.Getenv("DEV_DB_NAME")
        dbHost = os.Getenv("DEV_DB_HOST")
        dbPort = os.Getenv("DEV_DB_PORT")
    } else {
        pwd = os.Getenv("PROD_DB_PASSWORD")
        user = os.Getenv("PROD_DB_USER")
        dbName = os.Getenv("PROD_DB_NAME")
        dbHost = os.Getenv("PROD_DB_HOST")
        dbPort = os.Getenv("PROD_DB_PORT") }
	fmt.Println("Connecting to")
	fmt.Printf("%s:%s@tcp(%s:%s)/%s", user, pwd, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pwd, dbHost, dbPort, dbName))

	if err != nil {
		panic(err)
	}
	return db
}
