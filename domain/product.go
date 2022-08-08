package domain

type Product struct {
	ProductId          int32  `gorm:"primaryKey" json:"product_id"`
	ProductName        string `db:"product_name" json:"product_name"`
	CategoryId         int64  `db:"category_id" json:"category_id"`
	Price              int64  `db:"price" json:"price"`
	Stock              int64  `db:"stock" json:"stock"`
	ProductDescription string `db:"product_description" json:"product_description"`
}

// type Image struct {
// 	ProductId int32
// 	ImageUrl  string `db:"image_url"`
// }
