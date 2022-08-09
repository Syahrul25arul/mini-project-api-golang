package repostiory

import (
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/logger"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(client *gorm.DB) ProductRepositoryImpl {
	return ProductRepositoryImpl{client}
}

func (p ProductRepositoryImpl) SaveProduct(product *domain.Product) *errs.AppErr {

	// save product
	if result := p.db.Create(product); result.Error != nil {
		logger.Error("error insert data product " + result.Error.Error())
		return errs.NewUnexpectedError("error insert data product")
	}

	return nil
}

func (p ProductRepositoryImpl) GetAllProduct(page int) ([]domain.Product, *errs.AppErr) {
	// create slice for all product
	var products []domain.Product

	// set configuration pagination
	dataPerPage := 5
	// totalProduct := p.db.Find(&products).RowsAffected
	// totalPage := int(math.Ceil(float64(totalProduct)/float64(dataPerPage)))

	// jika page = 0, set page menjadi 1 untuk halaman pertama
	if page == 0 {
		page = 1
	}

	firstData := (dataPerPage * page) - dataPerPage

	// cek apakah ada error
	if result := p.db.Limit(dataPerPage).Offset(firstData).Find(&products); result.Error != nil {
		logger.Error("error get all data product from db " + result.Error.Error())
		return nil, errs.NewUnexpectedError("err get all data product from db")
	}
	return products, nil
}

func (p ProductRepositoryImpl) GetProductById(productId string) (*domain.Product, *errs.AppErr) {
	// buat variable untuk struct domain.Product
	var product domain.Product

	// cek apakah data ada atau tidak
	if result := p.db.First(&product, productId); result.Error != nil {
		logger.Error("error get data product by id not found " + result.Error.Error())
		return nil, errs.NewNotFoundError("product not found")
	}

	return &product, nil
}

func (p ProductRepositoryImpl) SetupProductDummy() {
	p.db.Exec("TRUNCATE TABLE products restart identity")
	products := []domain.Product{
		{
			ProductName:        "mie goreng sedap",
			CategoryId:         1,
			Price:              2500,
			Stock:              25,
			ProductDescription: "ini mie goreng sedap",
		},
		{
			ProductName:        "mie goreng udang sedap",
			CategoryId:         1,
			Price:              3200,
			Stock:              21,
			ProductDescription: "ini mie goreng udang sedap",
		},
		{
			ProductName:        "indomie kari ayam",
			CategoryId:         1,
			Price:              3000,
			Stock:              50,
			ProductDescription: "ini indomie kari ayam",
		},
		{
			ProductName:        "indomie goreng",
			CategoryId:         1,
			Price:              2700,
			Stock:              67,
			ProductDescription: "ini indomie goreng",
		},
		{
			ProductName:        "mie selera pedas",
			CategoryId:         1,
			Price:              2200,
			Stock:              31,
			ProductDescription: "ini mie selera pedas",
		},
		{
			ProductName:        "teh pucuk",
			CategoryId:         2,
			Price:              4000,
			Stock:              30,
			ProductDescription: "ini teh pucuk",
		},
		{
			ProductName:        "golda",
			CategoryId:         2,
			Price:              4000,
			Stock:              24,
			ProductDescription: "ini golda",
		},
		{
			ProductName:        "aqua",
			CategoryId:         2,
			Price:              5000,
			Stock:              26,
			ProductDescription: "ini aqua",
		},
		{
			ProductName:        "coca cola",
			CategoryId:         2,
			Price:              6000,
			Stock:              76,
			ProductDescription: "ini coca cola",
		},
		{
			ProductName:        "green tea",
			CategoryId:         2,
			Price:              8000,
			Stock:              38,
			ProductDescription: "ini green tea",
		},
		{
			ProductName:        "rokok sampoerna",
			CategoryId:         3,
			Price:              28000,
			Stock:              38,
			ProductDescription: "ini rokok sampoerna",
		},
		{
			ProductName:        "rokok surya pro",
			CategoryId:         3,
			Price:              18000,
			Stock:              27,
			ProductDescription: "ini surya pro",
		},
		{
			ProductName:        "rokok Malboro",
			CategoryId:         3,
			Price:              34000,
			Stock:              38,
			ProductDescription: "ini rokok Malboro",
		},
		{
			ProductName:        "Gula 1kg",
			CategoryId:         3,
			Price:              14000,
			Stock:              18,
			ProductDescription: "ini Gula 1 kg",
		},
		{
			ProductName:        "Garam",
			CategoryId:         3,
			Price:              4000,
			Stock:              58,
			ProductDescription: "ini Garam",
		},
	}
	p.db.Create(&products)
}
