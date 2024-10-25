package models

import "database/sql"

// Project represents the data structure for a project
type Project struct {
  ID          int            `json:"id"`
  Name        sql.NullString `json:"name"`
  Status      sql.NullString `json:"status"`
  Address     sql.NullString `json:"address"`
  City        sql.NullString `json:"city"`
  State       sql.NullString `json:"state"`
  Units       sql.NullInt32  `json:"units"`
  Stories     sql.NullString `json:"stories"`
  SqFt        sql.NullInt32  `json:"sq_ft"`
  Company_Id  sql.NullInt32  `json:"company_id"`
  Company     sql.NullString `json:"company"`
  Discipline  sql.NullString `json:"discipline"`
}