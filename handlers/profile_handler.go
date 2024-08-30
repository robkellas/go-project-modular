package handlers

import (
	"database/sql"
	"go-project-modular/models"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// GetProfileHandler returns a handler for getting a specific profile by ID
func GetProfileHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the profile ID from the URL
		vars := mux.Vars(r)
		profileID := vars["profileID"]

		// Initialize the Profile struct
		var profile models.Profile
		err := db.QueryRow("SELECT id, first_name, last_name, linkedin FROM profiles WHERE id = ? LIMIT 1", profileID).Scan(&profile.ID, &profile.FirstName, &profile.LastName, &profile.LinkedIn)
		if err != nil {
			// Handle the case where no profile is found
			if err == sql.ErrNoRows {
				http.NotFound(w, r)
				return
			}
			// Handle other database errors
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Parse the template files
		tmpl, err := template.ParseFiles(
			"templates/base.html",
			"templates/nav.html",
			"templates/footer.html",
			"templates/profile.html",
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the template with the profile data
		err = tmpl.ExecuteTemplate(w, "base.html", profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
