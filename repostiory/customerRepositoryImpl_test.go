package repostiory

import (
	"mini-project/config"
	"mini-project/database"
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/helper"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	godotenv.Load("./../.env")
	config.SanityCheck()
	// db := database.GetClientDb()
	// db.Exec("TRUNCATE TABLE users, customers restart identity")
	m.Run()
}

func TestCustomerRepositoryImpl_Register(t *testing.T) {
	// prepare database and repository
	db := database.GetClientDb()
	repository := NewCustomerRepository(db)

	testCase := []struct {
		name     string
		want     domain.Customer
		want2    domain.Users
		expected *errs.AppErr
	}{
		{
			name:     "Register success",
			want:     domain.Customer{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "inactive"},
			want2:    domain.Users{Username: "jamal12", Password: helper.BcryptPassword(config.SECRET_KEY + "password"), Role: "user"},
			expected: nil,
		},
		{
			name:     "Register failed field status empty struct customer",
			want:     domain.Customer{Name: "hendrik", DateOfBirth: "1995-02-25", ZipCode: "13421"},
			want2:    domain.Users{Username: "hendrik33", Password: helper.BcryptPassword(config.SECRET_KEY + "password"), Role: "user"},
			expected: errs.NewUnexpectedError("error insert data customer"),
		},
		{
			name:     "Register vailed field status is not active or inactive struct customer",
			want:     domain.Customer{Name: "array", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "test"},
			want2:    domain.Users{Username: "array97", Password: helper.BcryptPassword(config.SECRET_KEY + "password"), Role: "user"},
			expected: errs.NewUnexpectedError("error insert data customer"),
		},
		{
			name:     "Register vailed field duplicate key value violates struct customer",
			want:     domain.Customer{CustomerId: 1, Name: "rizal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active"},
			want2:    domain.Users{Username: "rizal91", Password: helper.BcryptPassword(config.SECRET_KEY + "password"), Role: "user"},
			expected: errs.NewUnexpectedError("error insert data customer"),
		},
		{
			name:     "Register failed field duplicate key value violates struct users",
			want:     domain.Customer{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active"},
			want2:    domain.Users{Username: "jamal12", Password: helper.BcryptPassword(config.SECRET_KEY + "password"), Role: "user"},
			expected: errs.NewUnexpectedError("error insert data user"),
		},
		{
			name:     "Register constraint check username not empty string struct users",
			want:     domain.Customer{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active"},
			want2:    domain.Users{Username: "", Password: helper.BcryptPassword(config.SECRET_KEY + "password"), Role: "user"},
			expected: errs.NewUnexpectedError("error insert data user"),
		},
		{
			name:     "Register constraint check password not empty string struct users",
			want:     domain.Customer{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active"},
			want2:    domain.Users{Username: "jamalo", Password: "", Role: "user"},
			expected: errs.NewUnexpectedError("error insert data user"),
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result := repository.Register(&testTable.want, &testTable.want2)
			assert.Equal(t, testTable.expected, result)
		})
	}
}
