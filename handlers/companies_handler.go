package handlers

import (
	"database/sql"
	"go-project-modular/models"
	"html/template"
	"log"
	"net/http"
)

// GetAllCompaniesHandler returns a handler for getting all companies
func GetAllCompaniesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Query to select all companies
		rows, err := db.Query(`SELECT
            id,
            name,
            discipline,
            headquarters,
            website
            FROM companies;`)
		if err != nil {
			log.Println("Error querying companies:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Slice to hold the companies
		var companies []models.Company

		// Loop through the rows and scan data into the slice
		for rows.Next() {
			var company models.Company
			err := rows.Scan(
				&company.ID,
				&company.Name,
				&company.Discipline,
				&company.Headquarters,
				&company.Website,
			)
			if err != nil {
				log.Println("Error scanning company row:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			companies = append(companies, company)
		}

		// Parse the template files (update if you have a companies.html template)
		tmpl, err := template.ParseFiles(
			"templates/base.html",
			"templates/nav.html",
			"templates/footer.html",
			"templates/companies.html", // Update if needed
		)
		if err != nil {
			log.Println("Error parsing templates:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Execute the template with the list of companies
		err = tmpl.ExecuteTemplate(w, "base.html", companies)
		if err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
