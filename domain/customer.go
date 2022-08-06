package domain

type Customer struct {
	CustomerId  int32  `gorm:"primaryKey"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	ZipCode     string `db:"zip_code"`
	Status      string `db:"status"`
}
