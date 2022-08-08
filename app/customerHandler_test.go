package app

import (
	"bytes"
	"encoding/json"
	"mini-project/config"
	"mini-project/database"
	"mini-project/domain"
	"mini-project/helper"
	"mini-project/repostiory"
	"mini-project/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	godotenv.Load("./../.env")
	config.SanityCheck()
	db := database.GetClientDb()
	db.Exec("TRUNCATE TABLE users, customers restart identity")
	db.Exec("TRUNCATE TABLE products restart identity")
	m.Run()
}

func TestCustomerHandler_RegisterCustomerHandler(t *testing.T) {
	// setup gin
	r := gin.Default()

	// set handler customer
	repo := repostiory.NewCustomerRepository(database.GetClientDb())
	service := service.NewCustomerService(repo)
	handler := CustomerHandler{Service: service}

	// set end point to testing
	r.POST("/register", handler.RegisterCustomerHandler)

	// set table testCase
	testCase := []struct {
		name            string
		request         domain.RegisterRequest
		expectedCode    int
		expectedMessage string
	}{
		{
			name:            "Register handler success",
			request:         domain.RegisterRequest{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active", Username: "jamal12", Password: "jamalPassword"},
			expectedCode:    201,
			expectedMessage: "registrasi berhasil",
		},
		{
			name:            "Register handler success with set default status inactive",
			request:         domain.RegisterRequest{Name: "petrus", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "test", Username: "petrus31", Password: "petrusPassword"},
			expectedCode:    201,
			expectedMessage: "registrasi berhasil",
		},
		{
			name:            "Register handler failed validation error field empty",
			request:         domain.RegisterRequest{Name: "hendrik", DateOfBirth: "1995-02-25", ZipCode: "13421", Username: "hendrik65", Password: "hendrik"},
			expectedCode:    http.StatusUnprocessableEntity,
			expectedMessage: "field Status cannot be empty",
		},
		{
			name:            "Register handler failed empty with status = ''",
			request:         domain.RegisterRequest{Name: "petrus", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: ""},
			expectedCode:    http.StatusUnprocessableEntity,
			expectedMessage: "field Status cannot be empty",
		},
		{
			name:            "Register handler failed empty with username = ''",
			request:         domain.RegisterRequest{Name: "test", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active", Password: "password"},
			expectedCode:    http.StatusUnprocessableEntity,
			expectedMessage: "field Username cannot be empty",
		},
		{
			name:            "Register service failed empty with password = ''",
			request:         domain.RegisterRequest{Name: "test", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active", Username: "test12"},
			expectedCode:    http.StatusUnprocessableEntity,
			expectedMessage: "field Password cannot be empty",
		},
		{
			name:            "Register Service failed duplicate primary key username",
			request:         domain.RegisterRequest{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active", Username: "jamal12", Password: "jamalPassword"},
			expectedCode:    http.StatusInternalServerError,
			expectedMessage: "error insert data user",
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			// set data request to bytes and put to NewRequest
			jsonValue, _ := json.Marshal(testTable.request)
			req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, testTable.expectedCode, w.Code)

			// get response body from handler
			response := w.Body.String()

			// clear double quote from response body
			response = helper.ClearDoubleCode(response)

			assert.Equal(t, testTable.expectedMessage, response)
		})
	}
}
