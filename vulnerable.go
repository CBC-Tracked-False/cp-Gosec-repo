package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/user", getUser)
	http.ListenAndServe(":8080", nil) // G114: no timeouts
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// ❌ User input directly used
	userID := r.URL.Query().Get("id")

	// ❌ SQL Injection (G201)
	connStr := "postgres://user:password@localhost/db?sslmode=disable"
	db, _ := sql.Open("postgres", connStr)
	query := fmt.Sprintf("SELECT * FROM users WHERE id = '%s'", userID)
	rows, _ := db.Query(query)
	defer rows.Close()

	// ❌ Command Injection (G204)
	cmd := exec.Command("sh", "-c", "echo "+userID)
	out, _ := cmd.Output()

	fmt.Fprintf(w, "Result: %s %v", out, rows)
}
