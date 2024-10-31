package handlers

import (
	"database/sql"
	"go-project-modular/models"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// GetProjectHandler returns a handler for getting a specific project by ID
func GetProjectHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the project ID from the URL
		vars := mux.Vars(r)
		projectID := vars["projectID"]

		// Initialize the Project struct
		var project models.Project
		err := db.QueryRow(`
      SELECT
        p.id,
        p."name",
        p.status,
        p.address,
        p.city,
        p."state",
        p.units,
        p.stories,
        p.sqft,
        c.id,
        c.name,
        c.discipline
      FROM projects AS p
      LEFT JOIN project_facts AS pf ON pf.project_id = p.id
      LEFT JOIN companies AS c ON c.id = pf.company_id
      WHERE p.id = ? LIMIT 1
    `, projectID).Scan(
			&project.ID,
			&project.Name,
			&project.Status,
			&project.Address,
			&project.City,
			&project.State,
			&project.Units,
			&project.Stories,
			&project.SqFt,
			&project.Company_Id,
			&project.Company,
			&project.Company_Discipline,
		)
		if err != nil {
			// Handle the case where no project is found
			if err == sql.ErrNoRows {
				http.NotFound(w, r)
				return
			}
			// Handle other database errors
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Fetch profiles for the project
		rows, err := db.Query(`
      SELECT
        p.id,
        p.first_name,
        p.last_name,
        p.linkedin
      FROM
        profiles AS p
      JOIN project_facts AS pf ON pf.people_id = p.id
      WHERE
        pf.project_id = ?
    `, project.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Populate the Profiles field
		profiles := []models.Profile{} // Use your Profile struct directly
		for rows.Next() {
			var profile models.Profile
			err := rows.Scan(
				&profile.ID,
				&profile.FirstName,
				&profile.LastName,
				&profile.LinkedIn, // Scan into the LinkedIn field of your Profile struct
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			profiles = append(profiles, profile)
		}

		// Parse the template files
		tmpl, err := template.ParseFiles(
			"templates/base.html",
			"templates/nav.html",
			"templates/footer.html",
			"templates/project.html",
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the template with the project data and profiles
		data := struct {
			models.Project
			Profiles []models.Profile `json:"profiles"` // Use your Profile struct here as well
		}{
			Project:  project,
			Profiles: profiles,
		}
		err = tmpl.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
