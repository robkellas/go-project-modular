package handlers

import (
	"database/sql"
	"go-project-modular/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetProjectHandler returns a handler for getting a specific project by ID
func GetProjectHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		projectID := vars["projectID"]

		var project models.Project
		err := db.QueryRow(`SELECT
			p.name,
			p.status,
			p.address,
			p.city,
			p.state,
			p.units,
			p.stories,
			p.sqft,
			c.name as company,
			c.discipline
			FROM projects AS p
			LEFT JOIN project_facts AS f ON f.project_id = p.id
			LEFT JOIN companies AS c ON c.id = f.company_id
			LEFT JOIN profiles AS pr ON pr.id = f.people_id
			WHERE p.id = ?
			GROUP BY p.name;`, projectID).Scan(
			&project.Name,
			&project.Status,
			&project.Address,
			&project.City,
			&project.State,
			&project.Units,
			&project.Stories,
			&project.SqFt,
			&project.Company,
			&project.Discipline,
		)

		if err != nil {
			log.Println("Error querying project:", err)
			if err == sql.ErrNoRows {
				http.NotFound(w, r)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Parse the template files, including the base, header, footer, and project templates
		tmpl, err := template.ParseFiles(
			"templates/base.html",
			"templates/nav.html",
			"templates/footer.html",
			"templates/project.html",
		)
		if err != nil {
			log.Println("Error parsing templates:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Log before attempting to write to the response
		log.Println("Executing template for project:", project.Name)

		// Execute the template with the project data
		if err = tmpl.ExecuteTemplate(w, "base.html", project); err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Log successful execution
		// log.Println("Successfully rendered project page for:", project.Name)
	}
}
