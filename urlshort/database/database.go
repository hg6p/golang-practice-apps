package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func OnInit() *sql.DB {

	connStr := "postgres://davor:1234@127.0.0.1:5432/urlshort?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Set the maximum number of open connections in the pool
	db.SetMaxOpenConns(5)

	// Set the maximum number of idle connections in the pool
	db.SetMaxIdleConns(2)

	// Set the maximum connection idle time (optional)
	db.SetConnMaxIdleTime(60 * time.Second)

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	return db
}

func GetUrls(db *sql.DB, path string) (string, error) {
	fmt.Println(path)
	var url string

	if err := db.QueryRow("SELECT url FROM url_data WHERE path = $1", path).Scan(&url); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("No url found")
		}
		return "", fmt.Errorf("Err")
	}
	return url, nil
}
