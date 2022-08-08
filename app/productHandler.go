package app

import (
	"mini-project/domain"
	"mini-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service service.ProductService
}

func (p ProductHandler) SaveProductHandler(ctx *gin.Context) {
	// tangkap request body dari client
	var product domain.Product
	ctx.ShouldBindJSON(&product)

	if err := p.Service.SaveProductService(product); err != nil {
		// jika terjdi error tampilkan error
		ctx.JSON(err.Code, err.Message)
	} else {
		// jika tidak error, berikan response ke client
		ctx.JSON(http.StatusCreated, map[string]any{
			"code":    http.StatusCreated,
			"status":  "ok",
			"message": "success create product",
		})
	}

}
