package domain

import "database/sql"

type Users struct {
	Username   string         `db:"username"`
	Password   string         `db:"password"`
	Salt       string         `db:"salt"`
	Role       string         `db:"role"`
	CustomerId sql.NullString `db:"customer_id"`
}
