package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"working-database/models"
)

type application struct {
	Models models.Models
}

func main() {

	// Move this to env. e.g export DB_CON="postgres://postgres:postgres@localhost/youtube?sslmode=disable"
	dsn := os.Getenv("DB_CON")

	db, err := connectToDb(dsn)
	if err != nil {
		log.Fatalln(err)
	}

	app := application{
		Models: models.NewModel(db),
	}

	fmt.Println("Starting application")
	err = app.serve()
	if err != nil {
		log.Fatalln(err)
	}
}

func connectToDb(dsn string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
