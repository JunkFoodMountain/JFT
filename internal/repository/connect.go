package repository

import (
	"database/sql"
	"fmt"
	"time"
)

func connectWithRetries(connStr string, maxRetries int, retryTimeout time.Duration) (*sql.DB, error) {
	var db *sql.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		// Attempt to connect to the database
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			fmt.Printf("Failed to connect to database on attempt %d: %s\n", i+1, err.Error())
			time.Sleep(retryTimeout)
			continue
		}

		// Ping the database to ensure a connection is made
		err = db.Ping()
		if err == nil {
			fmt.Println("Successfully connected to database!")
			return db, nil
		}

		// Close the database connection if it was made, but the ping failed
		closeErr := db.Close()
		if closeErr != nil {
			return nil, closeErr
		}

		fmt.Printf("Failed to ping database on attempt %d: %s\n", i+1, err.Error())
		time.Sleep(retryTimeout)
	}

	return nil, fmt.Errorf("could not connect to database after %d attempts", maxRetries)
}
