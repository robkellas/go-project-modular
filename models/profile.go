package models

import "database/sql"

// Profile represents the data structure for a profile
type Profile struct {
    ID        int
    FirstName string
    LastName  string
    LinkedIn  sql.NullString
}
