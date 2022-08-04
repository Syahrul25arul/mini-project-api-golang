package app

import (
	"fmt"
	"mini-project/config"
	"mini-project/logger"
	"net/http"

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

	// give info where server and port app running
	logger.Info(fmt.Sprintf("start server on  %s:%s ...", config.SERVER_ADDRESS, config.SERVER_PORT))

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	// run server
	r.Run(fmt.Sprintf("%s:%s", config.SERVER_ADDRESS, config.SERVER_PORT))

}
