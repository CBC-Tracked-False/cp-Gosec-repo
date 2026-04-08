package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")

	// ❌ Vulnerable: SQL Injection
	query := "SELECT * FROM users WHERE username = '" + username + "'"

	db, _ := sql.Open("mysql", "root:password@/testdb")
	rows, _ := db.Query(query)

	fmt.Fprintf(w, "Query executed: %v", rows)
}

func main() {
	http.HandleFunc("/user", handler)
	http.ListenAndServe(":8080", nil)
}
