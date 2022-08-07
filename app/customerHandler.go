package app

import (
	"mini-project/domain"
	"mini-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	Service service.CustomerService
}

func (handler CustomerHandler) RegisterCustomerHandler(c *gin.Context) {
	// tangkap request dari client
	var register domain.RegisterRequest
	c.ShouldBindJSON(&register)

	// jika terjadi error saat registrasi data customer, tampilkan error
	if err := handler.Service.RegisterCustomer(register); err != nil {
		c.JSON(err.Code, err.Message)
	} else {
		c.JSON(http.StatusCreated, "registrasi berhasil")
	}

}
