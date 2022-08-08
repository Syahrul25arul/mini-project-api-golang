package service

import (
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/response"
)

type AuthService interface {
	Login(request domain.LoginRequest) (*response.LoginResponse, *errs.AppErr)
}
