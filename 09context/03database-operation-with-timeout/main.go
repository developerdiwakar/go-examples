package main

import (
	"context"
	"fmt"
	"time"

	// Simulate database access (replace with actual database library)
	"database/sql"
)

func fetchData(ctx context.Context, db *sql.DB) (string, error) {
	var data string
	query := "SELECT data FROM my_table"
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second) // Set timeout
	defer cancel()                                         // Ensure cancellation on function exit

	row := db.QueryRowContext(ctx, query)
	err := row.Scan(&data)
	if err != nil {
		if err == context.DeadlineExceeded {
			return "", fmt.Errorf("database operation timed out")
		}
		return "", fmt.Errorf("error fetching data: %w", err)
	}
	return data, nil
}

func main() {
	// Simulate database connection (TODO: replace with actual connection logic)
	db, err := sql.Open("some_db", "user:password@tcp(localhost:3306)/my_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background() // Use background context here (no cancellation)
	data, err := fetchData(ctx, db)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Fetched data:", data)
}
