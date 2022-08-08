package app

import (
	"bytes"
	"encoding/json"
	"mini-project/database"
	"mini-project/domain"
	"mini-project/helper"
	"mini-project/repostiory"
	"mini-project/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProductHandler_SaveProductHandler(t *testing.T) {
	// setup gin
	r := gin.Default()

	// set handler product
	repo := repostiory.NewProductRepository(database.GetClientDb())
	service := service.NewProductService(repo)
	handler := ProductHandler{Service: service}

	// set end point to testing
	r.POST("/products", handler.SaveProductHandler)

	tests := []struct {
		name            string
		request         domain.Product
		expectedCode    int
		expectedMessage string
	}{
		// TODO: Add test cases.
		{
			name:            "product handler save product success",
			request:         domain.Product{ProductName: "teh pucuk", CategoryId: 2, Price: 10000, Stock: 20, ProductDescription: "ini teh pucuk"},
			expectedCode:    http.StatusCreated,
			expectedMessage: "{code:201,message:success create product,status:ok}",
		},
		{
			name:            "product handler save product description null success",
			request:         domain.Product{ProductName: "mie goreng", CategoryId: 1, Price: 5000, Stock: 20, ProductDescription: ""},
			expectedCode:    http.StatusCreated,
			expectedMessage: "{code:201,message:success create product,status:ok}",
		},
		{
			name:            "product handler save product failed duplicate primary key",
			request:         domain.Product{ProductId: 1, ProductName: "teh pucuk", CategoryId: 2, Price: 10000, Stock: 20, ProductDescription: "ini teh pucuk"},
			expectedCode:    http.StatusInternalServerError,
			expectedMessage: "error insert data product",
		},
		{
			name:            "product handler save product failed invalid categoryId not less than 1",
			request:         domain.Product{ProductName: "teh pucuk", CategoryId: 0, Price: 10000, Stock: 20, ProductDescription: "ini teh pucuk"},
			expectedCode:    http.StatusUnprocessableEntity,
			expectedMessage: "field CategoryId cannot less than 1",
		},
		{
			name:            "product handler save product failed invalid field price not less than 0",
			request:         domain.Product{ProductName: "teh pucuk", CategoryId: 2, Price: -1, Stock: 20, ProductDescription: "ini teh pucuk"},
			expectedCode:    http.StatusUnprocessableEntity,
			expectedMessage: "field Price cannot less than 0",
		},
		{
			name:            "product handler save product failed invalid field stock not less than 0",
			request:         domain.Product{ProductName: "teh pucuk", CategoryId: 2, Price: 4000, Stock: -1, ProductDescription: "ini teh pucuk"},
			expectedCode:    http.StatusUnprocessableEntity,
			expectedMessage: "field Stock cannot less than 0",
		},
		{
			name:            "product invalid field productName cannot be empty",
			request:         domain.Product{CategoryId: 2, Price: 4000, Stock: 20, ProductDescription: "ini teh pucuk"},
			expectedCode:    http.StatusUnprocessableEntity,
			expectedMessage: "field ProductName cannot be empty",
		},
	}
	for _, testTable := range tests {
		t.Run(testTable.name, func(t *testing.T) {
			// set data request to bytes and put to NewRequest
			jsonValue, _ := json.Marshal(testTable.request)
			req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonValue))

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
