package middleware

import (
	"fmt"
	"mini-project/config"
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/logger"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func isBearerToken(authorizationHeader string, c *gin.Context) bool {
	if !strings.Contains(authorizationHeader, "Bearer") {
		appErr := errs.NewBadRequestError("invalid token")
		c.JSON(appErr.Code, appErr)
		c.Abort()
		return false
	}
	return true
}

func AuthMiddleware() gin.HandlerFunc {
	// get Beearer token
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")

		// cek apakah token dari adalah bearer token
		if !isBearerToken(authorizationHeader, c) {
			return
		}

		// replace Bearer from bearer token
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		// parse token claims
		token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.Error("signing method invalid not !ok")
				return nil, fmt.Errorf("signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				logger.Error("signing method invalid method not signingmethodHS256")
				return nil, fmt.Errorf("signing method invalid")
			}

			return []byte(config.SECRET_KEY), nil
		})

		// cek if error from claims token
		if err != nil {
			appErr := errs.NewBadRequestError(err.Error())
			c.JSON(appErr.Code, appErr)
			c.Abort()
			return
		}

		// casting claims token to struct AccessTokenClaims
		_, ok := token.Claims.(*domain.AccessTokenClaims)

		// cek if token casting claims token not valid, return error
		if !ok || !token.Valid {
			appErr := errs.NewBadRequestError("invalid token")
			c.JSON(appErr.Code, appErr)
			c.Abort()
			return
		}

		c.Next()
	}
}

func IsAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")

		// cek apakah token dari adalah bearer token
		if !strings.Contains(authorizationHeader, "Bearer") {
			appErr := errs.NewBadRequestError("invalid token")
			c.JSON(appErr.Code, appErr)
			c.Abort()
			return
		}

		// replace Bearer from bearer token
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		// parse token claims
		token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.Error("signing method invalid not !ok")
				return nil, fmt.Errorf("signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				logger.Error("signing method invalid method not signingmethodHS256")
				return nil, fmt.Errorf("signing method invalid")
			}

			return []byte(config.SECRET_KEY), nil
		})

		// cek if error from claims token
		if err != nil {
			appErr := errs.NewBadRequestError(err.Error())
			c.JSON(appErr.Code, appErr)
			c.Abort()
			return
		}

		// casting claims token to struct AccessTokenClaims
		claims, ok := token.Claims.(*domain.AccessTokenClaims)

		// cek if token casting claims token not valid, return error
		if !ok || !token.Valid {
			appErr := errs.NewBadRequestError("invalid token")
			c.JSON(appErr.Code, appErr)
			c.Abort()
			return
		}

		// jika bukan admin, munculkan eror
		if claims.Role != "admin" {
			appErr := errs.NewForbiddenError("don't have enough permission to get the resources")
			c.JSON(appErr.Code, appErr)
			c.Abort()
			return
		}

		c.Next()
	}
}
