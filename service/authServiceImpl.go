package service

import (
	"fmt"
	"mini-project/config"
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/logger"
	"mini-project/repostiory"
	"mini-project/response"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repo repostiory.UserRepository
}

func NewAuthService(repo repostiory.UserRepository) AuthServiceImpl {
	return AuthServiceImpl{repo}
}

func (auth AuthServiceImpl) Login(request domain.LoginRequest) (*response.LoginResponse, *errs.AppErr) {
	// siapkan struct user dan error
	var err *errs.AppErr
	var user *domain.Users

	// ambil data user by username dari request
	// jika tidak ketemu, kembalikan error
	if user, err = auth.repo.FindByUsername(request.Username); err != nil {
		return nil, err
	}

	// cek apakah password sudah benar
	paswordSaltVerify := config.SECRET_KEY + request.Password
	if errorVerify := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(paswordSaltVerify)); errorVerify != nil {
		logger.Error(fmt.Sprintf("password from %s not verify", request.Username))
		return nil, errs.NewAuthenticationError("invalid credential")
	}

	// create claims token
	login := user.ToLogin()
	claims := login.ClaimsAccessToken()
	authToken := domain.NewAuthToken(claims)

	if accessToken, appErr := authToken.NewAccessToken(); appErr != nil {
		return nil, appErr
	} else {
		return response.NewLoginSucess(accessToken), nil
	}
}
