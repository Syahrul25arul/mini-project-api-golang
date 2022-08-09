package app

import (
	"fmt"
	"mini-project/config"
	"mini-project/database"
	"mini-project/logger"
	"mini-project/repostiory"
	"mini-project/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Start() {
	// loading env variabel
	if err := godotenv.Load(); err != nil {
		logger.Fatal("error loading file .env variable " + err.Error())
	}

	// check all variables are loaded
	config.SanityCheck()
	dbClient := database.GetClientDb()

	// prepare handler customer
	customerRepository := repostiory.NewCustomerRepository(dbClient)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := CustomerHandler{customerService}

	// prepare handle auth login
	userRepo := repostiory.NewUserRepository(dbClient)
	authService := service.NewAuthService(userRepo)
	authHandler := AuthHandler{authService}

	// prepare handle products
	productRepo := repostiory.NewProductRepository(dbClient)
	productService := service.NewProductService(productRepo)
	productHandler := ProductHandler{productService}

	// setup dummy product
	productRepo.SetupProductDummy()

	r := gin.Default()
	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "hello world",
	// 	})
	// })

	r.POST("/register", customerHandler.RegisterCustomerHandler)
	r.POST("/login", authHandler.LoginHandler)
	r.POST("/products", productHandler.SaveProductHandler)
	r.GET("/products", productHandler.GetAlProductHandler)
	r.GET("/products/:productId", productHandler.GetProdutById)

	// give info where server and port app running
	logger.Info(fmt.Sprintf("start server on  %s:%s ...", config.SERVER_ADDRESS, config.SERVER_PORT))

	// run server
	r.Run(fmt.Sprintf("%s:%s", config.SERVER_ADDRESS, config.SERVER_PORT))

}
