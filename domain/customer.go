package domain

type Customer struct {
	CustomerId  int32  `gorm:"primaryKey"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	ZipCode     string `db:"zip_code"`
	Status      string `db:"status"`
}

type RegisterRequest struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	ZipCode     string `json:"zip_code"`
	Status      string `json:"string"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

func (r RegisterRequest) ToCustomer() *Customer {
	return &Customer{
		Name:        r.Name,
		DateOfBirth: r.DateOfBirth,
		ZipCode:     r.ZipCode,
		Status:      r.Status,
	}
}

func (r RegisterRequest) ToUsers() *Users {
	return &Users{
		Username: r.Username,
		Password: r.Password,
	}
}
