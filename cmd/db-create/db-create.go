package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open the SQLite database file
	// The database file will be created if it doesn't exist
	db, err := sql.Open("sqlite3", "./profiles.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a table
	createTableSQL := `CREATE TABLE IF NOT EXISTS profiles (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"age" INTEGER
	  );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table created successfully")

	// Insert some data into the table
	insertprofilesQL := `INSERT INTO profiles (name, age) VALUES (?, ?)`
	_, err = db.Exec(insertprofilesQL, "Alice", 30)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}
	_, err = db.Exec(insertprofilesQL, "Bob", 25)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}
	fmt.Println("Data inserted successfully")

	// Query the data
	querySQL := `SELECT id, name, age FROM profiles`
	rows, err := db.Query(querySQL)
	if err != nil {
		log.Fatalf("Failed to query data: %v", err)
	}
	defer rows.Close()

	// Print the queried data
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Row iteration error: %v", err)
	}
}
