package models

import "database/sql"

// Project represents the data structure for a project
type Project struct {
    ID          int
    Name        sql.NullString
    Status      sql.NullString
    Address     sql.NullString
    City        sql.NullString
    State       sql.NullString
    Units       sql.NullInt32
    Stories     sql.NullString
    SqFt        sql.NullInt32
    Company_Id  sql.NullInt32
    Company     sql.NullString
    Discipline  sql.NullString
}
