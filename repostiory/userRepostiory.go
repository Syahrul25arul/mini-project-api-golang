package repostiory

import (
	"mini-project/domain"
	"mini-project/errs"
)

type UserRepository interface {
	SaveAdmin(user domain.Users) *errs.AppErr
}
