package domain

import "database/sql"

type Login struct {
	Username   string         `db:"username"`
	CustomerId sql.NullString `db:"customer_id"`
	Role       string         `db:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
