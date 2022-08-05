package repostiory

import (
	"mini-project/domain"
	"mini-project/errs"
)

type CustomerRepository interface {
	Register(customer domain.Customer) *errs.AppErr
}
