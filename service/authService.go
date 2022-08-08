package service

import (
	"mini-project/domain"
	"mini-project/errs"
)

type AuthService interface {
	Login(request domain.LoginRequest) (*domain.Users, *errs.AppErr)
}
