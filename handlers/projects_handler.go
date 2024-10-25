package handlers

import (
	"database/sql"
	"go-project-modular/models"
	"html/template"
	"log"
	"net/http"
)

// GetAllProjectsHandler returns a handler for getting all projects
func GetAllProjectsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Query to select all projects
		rows, err := db.Query(`SELECT
			p.id,
			p.name,
			p.status,
			p.address,
			p.city,
			p.state,
			p.units,
			p.stories,
			p.sqft,
			c.id as company_id,
			c.name as company,
			c.discipline
			FROM projects AS p
			LEFT JOIN project_facts AS f ON f.project_id = p.id
			LEFT JOIN companies AS c ON c.id = f.company_id
			LEFT JOIN profiles AS pr ON pr.id = f.people_id
			GROUP BY p.name
			ORDER BY p.id ASC
			;`)
		if err != nil {
			log.Println("Error querying projects:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Slice to hold the projects
		var projects []models.Project

		// Loop through the rows and scan data into the slice
		for rows.Next() {
			var project models.Project
			err := rows.Scan(
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
				log.Println("Error scanning project row:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			projects = append(projects, project)
		}

		// Parse the template files, including the base, header, footer, and projects templates
		tmpl, err := template.ParseFiles(
			"templates/base.html",
			"templates/nav.html",
			"templates/footer.html",
			"templates/projects.html",
		)
		if err != nil {
			log.Println("Error parsing templates:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Execute the template with the list of projects
		err = tmpl.ExecuteTemplate(w, "base.html", projects)
		if err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
