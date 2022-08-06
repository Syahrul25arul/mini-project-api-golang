package service

import (
	"fmt"
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/logger"
	"mini-project/repostiory"
	"reflect"
)

type CustomerServiceImpl struct {
	repo repostiory.CustomerRepository
}

func NewCustomerService(repo repostiory.CustomerRepository) CustomerServiceImpl {
	return CustomerServiceImpl{repo}
}

func (s CustomerServiceImpl) RegisterCustomer(customer domain.Customer) *errs.AppErr {
	// cek apakah data yang dimasukkan sudah valid
	if err := s.isValid(customer); err != nil {
		logger.Error("error invalid data customer " + err.Message)
		return err
	}

	// cek jika status != active and tidak kosong
	// maka set status ke nilai default yaitu inactive
	if customer.Status != "active" && customer.Status != "" {
		customer.Status = "inactive"
	}
	return s.repo.Register(customer)

}

func (s CustomerServiceImpl) isValid(customer domain.Customer) *errs.AppErr {
	// buat reflection untuk domain customer
	ref := reflect.TypeOf(customer)

	// cek apakah ada field customer yang belum di isi
	for i := 0; i < ref.NumField(); i++ {
		if ref.Field(i).Name != "CustomerId" && reflect.ValueOf(customer).Field(i).Interface() == "" {
			return errs.NewValidationError(fmt.Sprintf("field %s cannot be empty", ref.Field(i).Name))
		}
	}

	return nil
}
