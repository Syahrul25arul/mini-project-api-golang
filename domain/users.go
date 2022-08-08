package domain

import "database/sql"

type Users struct {
	Username   string        `db:"username"`
	Password   string        `db:"password"`
	Role       string        `db:"role"`
	CustomerId sql.NullInt32 `db:"customer_id"`
}

func (u *Users) SetRoleUser() {
	u.Role = "user"
}

func (u *Users) SetRoleAdmin() {
	u.Role = "admin"
}

func (u *Users) ToLogin() Login {
	return Login{Username: u.Username, CustomerId: sql.NullInt32{Int32: u.CustomerId.Int32, Valid: true}, Role: u.Role}
}
