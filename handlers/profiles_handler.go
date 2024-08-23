package handlers

import (
	"database/sql"
	"go-project-modular/models" // Import the models package
	"html/template"
	"net/http"
)

// GetAllProfilesHandler returns a handler for getting all profiles
func GetAllProfilesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, first_name, last_name, linkedin, archive FROM profiles")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var profiles []models.Profile

		for rows.Next() {
			var profile models.Profile
			var archive bool
			err := rows.Scan(&profile.ID, &profile.FirstName, &profile.LastName, &profile.LinkedIn, &archive)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if !archive {
				profiles = append(profiles, profile)
			}
		}

		err = rows.Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Parse the template files
		tmpl, err := template.ParseFiles(
			"templates/base.html",
			"templates/nav.html",
			"templates/footer.html", 
			"templates/profiles.html",
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the template with the profiles data
		err = tmpl.ExecuteTemplate(w, "base.html", profiles)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
