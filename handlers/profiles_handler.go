package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

// GetAllProfilesHandler returns a handler for getting all profiles
func GetAllProfilesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, age FROM profiles")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var name string
			var age int
			err := rows.Scan(&id, &name, &age)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "Profile ID: %d\nName: %s\nAge: %d\n\n", id, name, age)
		}

		err = rows.Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
