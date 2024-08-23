package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetProfileHandler returns a handler for getting a specific profile by ID
func GetProfileHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		profileID := vars["profileID"]

		var name string
		var age int
		err := db.QueryRow("SELECT name, age FROM profiles WHERE id = ?", profileID).Scan(&name, &age)
		if err != nil {
			if err == sql.ErrNoRows {
				http.NotFound(w, r)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Profile ID: %s\nName: %s\nAge: %d\n", profileID, name, age)
	}
}
