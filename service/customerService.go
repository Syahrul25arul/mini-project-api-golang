package service

import (
	"mini-project/domain"
	"mini-project/errs"
)

type CustomerService interface {
	RegisterCustomer(request domain.RegisterRequest) *errs.AppErr
}
