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

func (u UserRepositoryImpl) FindByUsername(username string) (*domain.Users, *errs.AppErr) {
	var user domain.Users
	if result := u.db.Where("username = ?", username).Find(&user); result.RowsAffected == 0 {
		logger.Error("error get data user by username not found")
		return nil, errs.NewNotFoundError("error get data user by username not found")
	}

	return &user, nil
}
