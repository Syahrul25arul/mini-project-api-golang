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

	// prepare handler customer
	dbClient := database.GetClientDb()
	customerRepository := repostiory.NewCustomerRepository(dbClient)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := CustomerHandler{customerService}

	r := gin.Default()
	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "hello world",
	// 	})
	// })

	r.POST("/register", customerHandler.RegisterCustomerHandler)

	// give info where server and port app running
	logger.Info(fmt.Sprintf("start server on  %s:%s ...", config.SERVER_ADDRESS, config.SERVER_PORT))

	// run server
	r.Run(fmt.Sprintf("%s:%s", config.SERVER_ADDRESS, config.SERVER_PORT))

}
