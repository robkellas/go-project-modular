package models

import "database/sql"

type Profile struct {
	ID        int            `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	LinkedIn  sql.NullString `json:"linkedin"`
	Projects  []Project      `json:"projects"`
}
