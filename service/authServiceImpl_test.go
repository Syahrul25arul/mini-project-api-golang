package service

import (
	"fmt"
	"mini-project/config"
	"mini-project/database"
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/helper"
	"mini-project/repostiory"
	"mini-project/response"
	"net/http"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupUsers() *gorm.DB {
	db := database.GetClientDb()
	repoCustomer := repostiory.NewCustomerRepository(db)
	customer := domain.Customer{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "inactive"}
	user := domain.Users{Username: "jamal12", Password: helper.BcryptPassword(config.SECRET_KEY + "password"), Role: "user"}

	repoCustomer.Register(&customer, &user)
	return db
}

func TestAuthServiceImpl_Login(t *testing.T) {
	// setup db and dummy data
	db := SetupUsers()

	// setup auth service
	userRepo := repostiory.NewUserRepository(db)
	service := NewAuthService(userRepo)
	request := domain.LoginRequest{Username: "jamal12", Password: "password"}

	resp, err := service.Login(request)
	fmt.Println(resp)
	fmt.Println(err)

	testCase := []struct {
		name      string
		want      domain.LoginRequest
		expected1 *response.LoginResponse
		expected2 *errs.AppErr
	}{
		{
			name:      "auth login success",
			want:      domain.LoginRequest{Username: "jamal12", Password: "password"},
			expected1: &response.LoginResponse{Message: "success login", Code: http.StatusOK},
			expected2: nil,
		}, {
			name:      "auth login failed user not found",
			want:      domain.LoginRequest{Username: "notfound", Password: "notfound"},
			expected1: nil,
			expected2: errs.NewAuthenticationError("invalid credential"),
		},
		{
			name:      "auth login failed user password wrong",
			want:      domain.LoginRequest{Username: "jamal12", Password: "notfound"},
			expected1: nil,
			expected2: errs.NewAuthenticationError("invalid credential"),
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			resultResponse, resultErr := service.Login(testTable.want)

			// cek jika data response ada
			if resultResponse != nil {
				assert.Equal(t, testTable.expected1.Code, resultResponse.Code)
				assert.Equal(t, testTable.expected1.Message, resultResponse.Message)

				// cek jwt valid atau tiidak
				token, err := jwt.Parse(resultResponse.Token, func(token *jwt.Token) (interface{}, error) {
					return []byte(config.SECRET_KEY), nil
				})

				assert.True(t, token.Valid)
				assert.Nil(t, err)

			}

			assert.Equal(t, testTable.expected2, resultErr)

		})
	}
}
