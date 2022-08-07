package repostiory

import (
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/logger"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(client *gorm.DB) UserRepositoryImpl {
	return UserRepositoryImpl{client}
}

func (u UserRepositoryImpl) SaveUser(user *domain.Users) *errs.AppErr {
	if result := u.db.Create(user); result.Error != nil {
		logger.Error("error insert data user : " + result.Error.Error())
		return errs.NewUnexpectedError("error insert data user")
	}

	return nil
}
