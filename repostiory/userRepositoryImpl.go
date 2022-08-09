package repostiory

import (
	"mini-project/config"
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/helper"
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
		return nil, errs.NewAuthenticationError("invalid credential")
	}

	return &user, nil
}

func (u UserRepositoryImpl) SetupAdminDummy() {
	u.db.Exec("TRUNCATE TABLE users, customers restart identity")
	admin := domain.Users{
		Username: "admin",
		Password: helper.BcryptPassword(config.SECRET_KEY + "admin"),
		Role:     "admin",
	}
	u.db.Create(admin)
}
