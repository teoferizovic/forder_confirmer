package model

import "database/sql"

type Product struct {
	ID 	 sql.NullInt64              `json:"id"`
	Name sql.NullString             `json:"name"`
}
