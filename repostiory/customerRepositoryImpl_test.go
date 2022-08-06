package repostiory

import (
	"mini-project/config"
	"mini-project/database"
	"mini-project/domain"
	"mini-project/errs"
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

func TestCustomerRepositoryImpl_Register(t *testing.T) {
	// prepare database and repository
	db := database.GetClientDb()
	repository := NewCustomerRepository(db)

	testCase := []struct {
		name     string
		want     domain.Customer
		expected *errs.AppErr
	}{
		{
			name:     "Register success",
			want:     domain.Customer{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "inactive"},
			expected: nil,
		},
		{
			name:     "Register vailed field status empty",
			want:     domain.Customer{Name: "hendrik", DateOfBirth: "1995-02-25", ZipCode: "13421"},
			expected: errs.NewUnexpectedError("error insert data customer"),
		},
		{
			name:     "Register vailed field status is not active or inactive",
			want:     domain.Customer{Name: "array", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "test"},
			expected: errs.NewUnexpectedError("error insert data customer"),
		},
		{
			name:     "Register vailed field duplicate key value violates",
			want:     domain.Customer{CustomerId: 1, Name: "rizal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active"},
			expected: errs.NewUnexpectedError("error insert data customer"),
		},
		{
			name:     "Register vailed field duplicate key value violates",
			want:     domain.Customer{Name: "array", DateOfBirth: "1995-02-25", ZipCode: "13421"},
			expected: errs.NewUnexpectedError("error insert data customer"),
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result := repository.Register(testTable.want)
			assert.Equal(t, testTable.expected, result)
		})
	}
}
