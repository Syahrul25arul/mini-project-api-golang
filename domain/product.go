package domain

type Product struct {
	ProductId          int32  `gorm:"primaryKey"`
	ProductName        string `db:"product_name"`
	CategoryId         int64  `db:"category_id"`
	Price              int64  `db:"price"`
	Stock              int64  `db:"stock"`
	ProductDescription string `db:"product_description"`
}

// type Image struct {
// 	ProductId int32
// 	ImageUrl  string `db:"image_url"`
// }
