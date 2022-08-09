package repostiory

import (
	"mini-project/domain"
	"mini-project/errs"
)

type ProductRepository interface {
	SaveProduct(product *domain.Product) *errs.AppErr
	GetAllProduct(page int) ([]domain.Product, *errs.AppErr)
	GetProductById(productId string) (*domain.Product, *errs.AppErr)
}
