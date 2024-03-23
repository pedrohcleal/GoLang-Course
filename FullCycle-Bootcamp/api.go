package main

import {
	"net/http"
	"encoding/json"
	"net/http"
}

type user struct {
	ID int
	Name string
	Email string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc()
	http.ListenAndServe(":8080",nil)
}

func listUsersHandler( w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "go.db")
	if err !- nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
}


