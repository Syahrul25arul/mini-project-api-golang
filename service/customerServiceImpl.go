package service

import (
	"fmt"
	"mini-project/config"
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/helper"
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

func (s CustomerServiceImpl) RegisterCustomer(request domain.RegisterRequest) *errs.AppErr {
	// cek apakah data yang dimasukkan sudah valid
	if err := s.isValid(request); err != nil {
		logger.Error("error invalid data customer " + err.Message)
		return err
	}

	// ubah data register ke user dan customer struct
	customer := request.ToCustomer()
	user := request.ToUsers()

	// cek jika status != active and tidak kosong
	// maka set status ke nilai default yaitu inactive
	if customer.Status != "active" && customer.Status != "" {
		customer.Status = "inactive"
	}

	// hashing password sebelum di insert ke database dan set role jadi user
	user.Password = helper.BcryptPassword(config.SECRET_KEY + request.Password)
	user.SetRoleUser()

	return s.repo.Register(customer, user)

}

func (s CustomerServiceImpl) isValid(register domain.RegisterRequest) *errs.AppErr {
	// buat reflection untuk domain register
	ref := reflect.TypeOf(register)

	// cek apakah ada field register yang belum di isi
	for i := 0; i < ref.NumField(); i++ {
		if reflect.ValueOf(register).Field(i).Interface() == "" {
			return errs.NewValidationError(fmt.Sprintf("field %s cannot be empty", ref.Field(i).Name))
		}
	}

	return nil
}
