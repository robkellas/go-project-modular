package main

import (
	"fmt"
	"log"
	"net/http"

	// Import the db package
	// Import the handlers package
	"go-project-modular/db"
	"go-project-modular/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Set up the router and routes and config the trailing slash to redirect to no slash
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", handlers.HomeHandler()).Methods("GET")
	r.HandleFunc("/profiles/{profileID}/", handlers.GetProfileHandler(database)).Methods("GET")
	r.HandleFunc("/profiles/", handlers.GetAllProfilesHandler(database)).Methods("GET")

	// Start the server on port 8080
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
