package service

import (
	"mini-project/domain"
	"mini-project/errs"
)

type ProductService interface {
	SaveProductService(product domain.Product) *errs.AppErr
	GetAllProductService(page int) ([]domain.Product, *errs.AppErr)
	GetProductByIdService(productId string) (*domain.Product, *errs.AppErr)
	DeleteProductService(productId string) *errs.AppErr
}
