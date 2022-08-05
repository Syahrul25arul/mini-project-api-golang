package repostiory

import (
	"mini-project/config"
	"mini-project/database"
	"mini-project/domain"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepositoryImpl_Register(t *testing.T) {
	godotenv.Load("./../.env")

	config.SanityCheck()

	db := database.GetClientDb()
	repository := NewCustomerRepository(db)

	customer := domain.Customer{CustomerId: 2, Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "inactive"}

	err := repository.Register(customer)

	assert.Nil(t, err)
}
