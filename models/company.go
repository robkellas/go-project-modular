package models

import "database/sql"
type Company struct {
  ID                    int            `json:"id"`
  Name                  sql.NullString `json:"name"`
  Discipline            sql.NullString `json:"discipline"`
  Headquarters          sql.NullString `json:"headquarters"`
  Website               sql.NullString `json:"website"`
}