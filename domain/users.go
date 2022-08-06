package domain

import "database/sql"

type Users struct {
	Username   string        `db:"username"`
	Password   string        `db:"password"`
	Role       string        `db:"role"`
	CustomerId sql.NullInt32 `db:"customer_id"`
}
