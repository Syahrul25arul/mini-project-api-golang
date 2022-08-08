package app

import (
	"bytes"
	"encoding/json"
	"mini-project/config"
	"mini-project/database"
	"mini-project/domain"
	"mini-project/helper"
	"mini-project/repostiory"
	"mini-project/response"
	"mini-project/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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

func TestAuthHandler_LoginHandler(t *testing.T) {
	// setup gin
	r := gin.Default()

	// set handler auth
	repo := repostiory.NewUserRepository(SetupUsers())
	service := service.NewAuthService(repo)
	handler := AuthHandler{Service: service}

	// set end point to testing
	r.POST("/login", handler.LoginHandler)

	testCase := []struct {
		name      string
		request   domain.LoginRequest
		expected1 int
		expected2 string
	}{
		{
			name:      "auth login success",
			request:   domain.LoginRequest{Username: "jamal12", Password: "password"},
			expected1: http.StatusOK,
			expected2: "success login",
		}, {
			name:      "auth login failed user not found",
			request:   domain.LoginRequest{Username: "notfound", Password: "notfound"},
			expected1: http.StatusUnauthorized,
			expected2: "invalid credential",
		},
		{
			name:      "auth login failed user password wrong",
			request:   domain.LoginRequest{Username: "jamal12", Password: "notfound"},
			expected1: http.StatusUnauthorized,
			expected2: "invalid credential",
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			// set data request to bytes and put to NewRequest
			jsonValue, _ := json.Marshal(testTable.request)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if testTable.expected1 == http.StatusOK {
				// get response body from handler
				var response response.LoginResponse
				body := w.Body.String()
				json.Unmarshal([]byte(body), &response)

				assert.Equal(t, testTable.expected1, response.Code)
				assert.Equal(t, testTable.expected2, response.Message)
			} else {
				response := w.Body.String()

				// clear double code
				response = helper.ClearDoubleCode(response)

				assert.Equal(t, testTable.expected1, w.Code)
				assert.Equal(t, testTable.expected2, response)
			}

		})
	}
}
