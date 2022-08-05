package repostiory

import (
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/logger"

	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(client *gorm.DB) CustomerRepositoryImpl {
	return CustomerRepositoryImpl{client}
}

func (c CustomerRepositoryImpl) Register(customer domain.Customer) *errs.AppErr {
	if result := c.db.Create(&customer); result.Error != nil {
		logger.Error("error insert dat customer " + result.Error.Error())
		return errs.NewUnexpectedError("error insert data customer")
	}

	logger.Info("insert data customer success")
	return nil

}
