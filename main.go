package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

// Define a struct to parse the JSON request body
type RequestBody struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Keytext string `json:"keytext"`
}

func main() {
	// Initialize SQLite database
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if it doesn't exist
	createTableSQL := `CREATE TABLE IF NOT EXISTS data (
        "id" TEXT PRIMARY KEY,
        "email" TEXT NOT NULL,
        "keytext" TEXT NOT NULL,
        "timestamp" TEXT NOT NULL
    );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Define the handler for the /api/put endpoint
	http.HandleFunc("/api/put", func(w http.ResponseWriter, r *http.Request) {
		// Only allow POST requests
		if r.Method != "POST" {
			http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
			return
		}

		// Parse the JSON body of the request
		var body RequestBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		timestamp := time.Now().Format(time.RFC3339)

		// Print the values with the current timestamp
		fmt.Printf("Timestamp: %s, ID: %s, Email: %s, Keytext: %s\n", timestamp, body.ID, body.Email, body.Keytext)

		// Insert data into SQLite database
		insertSQL := `INSERT INTO data(id, email, keytext, timestamp) VALUES (?, ?, ?, ?)`
		statement, err := db.Prepare(insertSQL)
		if err != nil {
			// log.Panic(err)
			// log.Fatal(err)
			fmt.Println("Something went wrong while preparing for the SQL.")
		}
		_, err = statement.Exec(body.ID, body.Email, body.Keytext, timestamp)
		if err != nil {
			// fmt.Println("failed during execution")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error occured.")
			// log.Fatal(err)
		} else {
			// Respond to the client
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Data stored successfully: %s, %s, %s", body.ID, body.Email, body.Keytext)
		}

	})

	// Start the server on port 8080
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
