// company_handler.go

package handlers

import (
	"database/sql"
	"go-project-modular/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetCompanyHandler returns a handler for getting a single company and its associated projects
func GetCompanyHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the company ID from the URL
		vars := mux.Vars(r)
		companyID := vars["companyID"]
		// Fetch company details
		var company models.Company
		err := db.QueryRow(`SELECT 
   id, 
   name, 
   discipline, 
   headquarters, 
   website 
   FROM companies 
   WHERE id = ?`, companyID).Scan(
			&company.ID,
			&company.Name,
			&company.Discipline,
			&company.Headquarters,
			&company.Website,
		)
		if err != nil {
			log.Println("Error querying company:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Fetch projects associated with the company
		rows, err := db.Query(`SELECT DISTINCT
			p.id, 
			p.name,
			p.status,
			p.city,
			p.state,
			p.units,
			p.stories,
			p.sqft 
			FROM projects AS p 
			JOIN project_facts AS pf ON p.id = pf.project_id 
			WHERE pf.company_id = ?`, companyID)
		if err != nil {
			log.Println("Error querying projects:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var projects []models.Project
		for rows.Next() {
			var project models.Project
			err := rows.Scan(
				&project.ID,
				&project.Name,
				&project.Status,
				&project.City,
				&project.State,
				&project.Units,
				&project.Stories,
				&project.SqFt,
			)
			if err != nil {
				log.Println("Error scanning project:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			projects = append(projects, project)
		}

		// Prepare data for the template
		data := struct {
			Company  models.Company
			Projects []models.Project
		}{
			Company:  company,
			Projects: projects,
		}

		// Parse the template
		tmpl, err := template.ParseFiles(
			"templates/base.html",
			"templates/nav.html",
			"templates/footer.html",
			"templates/company.html", // Assuming you'll create this template
		)
		if err != nil {
			log.Println("Error parsing templates:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Execute the template
		err = tmpl.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
