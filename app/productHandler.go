package app

import (
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/logger"
	"mini-project/service"
	"net/http"
	"strconv"

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

func (p ProductHandler) GetAlProductHandler(ctx *gin.Context) {
	// tangkap query string dari client
	queryString := ctx.Query("page")

	var page int
	var err error

	// cek apakah query string page tidak kosong
	// jika kosong, biarkan ke nilai default yaitu 0
	if queryString != "" {
		// jika tidak kosong, casting ke int
		page, err = strconv.Atoi(queryString)

		if err != nil {
			logger.Error("error casting query string " + err.Error())
			ctx.JSON(http.StatusBadRequest, errs.NewBadRequestError("url not valid"))
			return
		}
	}

	// get data products dari db
	if products, errProduct := p.Service.GetAllProductService(page); errProduct != nil {
		// jika ada error tampilkan eror
		ctx.JSON(errProduct.Code, errProduct)
		return
	} else {
		ctx.JSON(http.StatusOK, products)
	}
}

func (p ProductHandler) GetProdutById(ctx *gin.Context) {
	// tangkap parameter productId dari url request
	productId := ctx.Param("productId")

	// cek jika parameter id adalah empty string, kembalikan pesan errro
	// karna akan terjadi error jika get data pada database dengan product id empty string
	if productId == "" {
		logger.Error("product id = empty string")
		ctx.JSON(http.StatusBadRequest, errs.NewBadRequestError("invalid url, parameter id not valid"))
		return
	}

	if product, err := p.Service.GetProductByIdService(productId); err != nil {
		ctx.JSON(err.Code, err.Message)
		return
	} else {
		ctx.JSON(http.StatusOK, product)
	}

}
