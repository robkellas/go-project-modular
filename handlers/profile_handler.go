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

    // Fetch projects for the profile
    rows, err := db.Query(`
      SELECT
				pr.id,
        pr."name",
				c.id,
        c."name",
        pr.city,
        pr."state",
        pr.units,
        pr.sqft,
        pr.stories
      FROM
        profiles AS p
      JOIN project_facts AS pf ON pf.people_id = p.id
      JOIN projects AS pr ON pr.id = pf.project_id
      JOIN companies AS c ON c.id = pf.company_id
      WHERE
        p.id = ?
    `, profile.ID)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    defer rows.Close()

    // Populate the Projects field
    profile.Projects = []models.Project{} // Initialize the slice
    for rows.Next() {
      var project models.Project
      err := rows.Scan(
        &project.ID,
				&project.Name,
				&project.Company_Id,
        &project.Company,
        &project.City,
        &project.State,
        &project.Units,
        &project.SqFt,
        &project.Stories,
      )
      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }
      profile.Projects = append(profile.Projects, project)
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
