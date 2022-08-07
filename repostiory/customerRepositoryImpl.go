package repostiory

import (
	"database/sql"
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

func (c CustomerRepositoryImpl) Register(customer *domain.Customer, user *domain.Users) *errs.AppErr {
	// begin transcation
	tx := c.db.Begin()

	// cek apakah saat insert data customer ada error
	// jika error rollback
	if resultCreateCustomer := tx.Create(customer); resultCreateCustomer.Error != nil {
		tx.Rollback()
		logger.Error("error insert dat customer " + resultCreateCustomer.Error.Error())
		return errs.NewUnexpectedError("error insert data customer")
	}

	// ambil last id customer dan masukkan kedalam struct user
	user.CustomerId = sql.NullInt32{Int32: customer.CustomerId, Valid: true}

	// cek apakah saat insert data customer ada error
	// jika error rollback
	repoUser := NewUserRepository(tx)
	if err := repoUser.SaveUser(user); err != nil {
		tx.Rollback()
		return err
	}

	// jika data customer dan akun user dari customer berhasil disimpan. maka commit
	tx.Commit()

	logger.Info("register data customer success")
	return nil

}
