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
	db.Exec("TRUNCATE TABLE users, customers, products restart identity")
	m.Run()
}

func TestCustomerServiceImpl_RegisterCustomer(t *testing.T) {
	// prepare database, repo and service
	db := database.GetClientDb()
	repo := repostiory.NewCustomerRepository(db)
	service := NewCustomerService(repo)

	testCase := []struct {
		name     string
		want     domain.RegisterRequest
		expected *errs.AppErr
	}{
		{
			name:     "Register Service success",
			want:     domain.RegisterRequest{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active", Username: "jamal12", Password: "jamalPassword"},
			expected: nil,
		},
		{
			name:     "Register service success with set default status inactive",
			want:     domain.RegisterRequest{Name: "petrus", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "test", Username: "petrus31", Password: "petrusPassword"},
			expected: nil,
		},
		{
			name:     "Register Service failed validation error field empty",
			want:     domain.RegisterRequest{Name: "hendrik", DateOfBirth: "1995-02-25", ZipCode: "13421", Username: "hendrik65", Password: "hendrik"},
			expected: errs.NewValidationError("field Status cannot be empty"),
		},
		{
			name:     "Register service failed empty with status = ''",
			want:     domain.RegisterRequest{Name: "petrus", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: ""},
			expected: errs.NewValidationError("field Status cannot be empty"),
		},
		{
			name:     "Register service failed empty with username = ''",
			want:     domain.RegisterRequest{Name: "test", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active", Password: "password"},
			expected: errs.NewValidationError("field Username cannot be empty"),
		},
		{
			name:     "Register service failed empty with password = ''",
			want:     domain.RegisterRequest{Name: "test", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active", Username: "test12"},
			expected: errs.NewValidationError("field Password cannot be empty"),
		},
		{
			name:     "Register Service failed duplicate primary key username",
			want:     domain.RegisterRequest{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active", Username: "jamal12", Password: "jamalPassword"},
			expected: errs.NewUnexpectedError("error insert data user"),
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result := service.RegisterCustomer(testTable.want)
			assert.Equal(t, testTable.expected, result)
		})
	}
}
