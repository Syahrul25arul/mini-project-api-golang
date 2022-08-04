package repostiory

import (
	"mini-project/domain"
	"mini-project/errs"
)

type userRepository interface {
	SaveAdmin(user domain.Users) *errs.AppErr
}
