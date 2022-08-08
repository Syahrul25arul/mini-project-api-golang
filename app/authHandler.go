package app

import (
	"mini-project/domain"
	"mini-project/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service service.AuthService
}

func (au AuthHandler) LoginHandler(c *gin.Context) {
	// tangkap login request
	var login domain.LoginRequest
	c.ShouldBindJSON(&login)

	// jika terjadi error saat registrasi, tampilkan error
	if resp, err := au.Service.Login(login); err != nil {
		c.JSON(err.Code, err.Message)
	} else {
		c.JSON(resp.Code, resp)
	}
}
