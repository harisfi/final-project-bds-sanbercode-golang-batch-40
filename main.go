package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	database "github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Environment Failed to Load")
		panic(err)
	} else {
		fmt.Println("Environment Successfully Loaded")
	}

	port, _ := strconv.Atoi(os.Getenv("PGPORT"))
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		port,
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()

	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)
	defer DB.Close()
}
