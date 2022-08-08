package repostiory

import (
	"mini-project/domain"
	"mini-project/errs"
)

type ProductRepository interface {
	SaveProduct(product *domain.Product) *errs.AppErr
}
