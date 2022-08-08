package service

import (
	"mini-project/domain"
	"mini-project/errs"
)

type ProductService interface {
	SaveProductService(product domain.Product) *errs.AppErr
}
