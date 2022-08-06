package service

import (
	"mini-project/domain"
	"mini-project/errs"
)

type CustomerService interface {
	RegisterCustomer(customer domain.Customer) *errs.AppErr
}
