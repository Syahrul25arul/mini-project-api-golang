package repostiory

import (
	"mini-project/domain"
	"mini-project/errs"
)

type UserRepository interface {
	// SaveAdmin(user domain.Users) *errs.AppErr
	SaveUser(user *domain.Users) *errs.AppErr
	FindByUsername(username string) (*domain.Users, *errs.AppErr)
}
