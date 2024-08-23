package handlers

import (
	"html/template"
	"net/http"
)

// HomeHandler returns a handler for rendering the homepage
func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the template files, including the base, header, footer, and homepage templates
		tmpl, err := template.ParseFiles(
			"templates/base.html",
			"templates/nav.html",
			"templates/footer.html",
			"templates/home.html",
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the template without any specific data
		err = tmpl.ExecuteTemplate(w, "base.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
