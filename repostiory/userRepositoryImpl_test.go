package repostiory

import (
	"database/sql"
	"mini-project/config"
	"mini-project/database"
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/helper"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupUsers() *gorm.DB {
	db := database.GetClientDb()
	repoCustomer := NewCustomerRepository(db)
	customer := domain.Customer{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "inactive"}
	user := domain.Users{Username: "jamal12", Password: helper.BcryptPassword(config.SECRET_KEY + "password"), Role: "user"}

	repoCustomer.Register(&customer, &user)
	return db
}

func TestGetAllUser(t *testing.T) {
	db := SetupUsers()
	repo := NewUserRepository(db)

	testCase := []struct {
		name      string
		want      string
		expected1 *domain.Users
		expected2 *errs.AppErr
	}{
		{
			name:      "find user success",
			want:      "jamal12",
			expected1: &domain.Users{Username: "jamal12", Role: "user", CustomerId: sql.NullInt32{Int32: 1, Valid: true}},
			expected2: nil,
		},
		{
			name:      "find user failed not found",
			want:      "test",
			expected1: nil,
			expected2: errs.NewAuthenticationError("invalid credential"),
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result, err := repo.FindByUsername(testTable.want)

			if err == nil {
				assert.Equal(t, result.Username, testTable.expected1.Username)
				assert.Equal(t, result.CustomerId, testTable.expected1.CustomerId)
				assert.Equal(t, result.Role, testTable.expected1.Role)
			}
			assert.Equal(t, err, testTable.expected2)
		})
	}
}
