package cmd

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	// DB The database connection
	DbConn *sqlx.DB
)

func Connect() (*sqlx.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbServer := os.Getenv("DB_SERVER")

	// Generating databaseUrl & not fetching the URL from the .env file
	databaseUrl := "postgres://" + dbUser + ":" + dbPass + "@" + dbServer + "/" + dbName
	db, err := sqlx.Connect("postgres", databaseUrl)

	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return db, nil
}
