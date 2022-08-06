package service

import (
	"mini-project/config"
	"mini-project/database"
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/repostiory"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	godotenv.Load("./../.env")
	config.SanityCheck()
	db := database.GetClientDb()
	db.Exec("TRUNCATE TABLE users, customers restart identity")
	m.Run()
}

func TestCustomerServiceImpl_RegisterCustomer(t *testing.T) {
	// prepare database, repo and service
	db := database.GetClientDb()
	repo := repostiory.NewCustomerRepository(db)
	service := NewCustomerService(repo)

	testCase := []struct {
		name     string
		want     domain.Customer
		expected *errs.AppErr
	}{
		{
			name:     "Register Service success",
			want:     domain.Customer{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active"},
			expected: nil,
		},
		{
			name:     "Register service success with set default status inactive",
			want:     domain.Customer{Name: "petrus", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "test"},
			expected: nil,
		},
		{
			name:     "Register Service failed validation error field empty",
			want:     domain.Customer{Name: "hendrik", DateOfBirth: "1995-02-25", ZipCode: "13421"},
			expected: errs.NewValidationError("field Status cannot be empty"),
		},
		{
			name:     "Register service failed empty with status = ''",
			want:     domain.Customer{Name: "petrus", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: ""},
			expected: errs.NewValidationError("field Status cannot be empty"),
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result := service.RegisterCustomer(testTable.want)
			assert.Equal(t, testTable.expected, result)
		})
	}
}
