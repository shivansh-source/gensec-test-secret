package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Blank import for SQL driver
)

func main() {
	// VULNERABILITY 1: Hardcoded Secret (Found by Gitleaks/Trivy)
	// The database password is exposed in the source code.
	db, err := sql.Open("mysql", "admin:Pa$$w0rd123@tcp(127.0.0.1:3306)/users_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("id")

		// Fix SQL Injection vulnerability by using parameterized queries
		query := "SELECT username FROM users WHERE user_id = ?"
		
		var username string
		err := db.QueryRow(query, userID).Scan(&username)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, "Username: %s", username)
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}