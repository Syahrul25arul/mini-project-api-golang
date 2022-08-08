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
