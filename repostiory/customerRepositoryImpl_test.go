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
	db.Exec("DELETE FROM customers")
	m.Run()
}

func TestCustomerRepositoryImpl_Register(t *testing.T) {
	db := database.GetClientDb()
	repository := NewCustomerRepository(db)

	// customer := domain.Customer{CustomerId: 1, Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "inactive"}

	// err := repository.Register(customer)

	// assert.Nil(t, err)

	testCase := []struct {
		name     string
		want     domain.Customer
		expected *errs.AppErr
	}{
		{
			name:     "Register success",
			want:     domain.Customer{CustomerId: 1, Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "inactive"},
			expected: nil,
		},
		{
			name:     "Register success",
			want:     domain.Customer{CustomerId: 2, Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421"},
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
