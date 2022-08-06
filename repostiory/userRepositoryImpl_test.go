package repostiory

import (
	"database/sql"
	"mini-project/config"
	"mini-project/database"
	"mini-project/domain"
	"mini-project/errs"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func bcryptPassword(passwordSalt string) string {
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordSalt), 8)
	return string(newPassword)
}

func TestUserRepositoryImpl_SaveUser(t *testing.T) {
	// prepare database and repository
	db := database.GetClientDb()
	repository := NewUserRepository(db)

	customerRepo := NewCustomerRepository(db)
	newCustomer := domain.Customer{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active"}
	customerRepo.Register(newCustomer)

	testCase := []struct {
		name     string
		want     domain.Users
		expected *errs.AppErr
	}{
		{
			name:     "Save user success",
			want:     domain.Users{Username: "hendrik", Password: bcryptPassword(config.SECRET_KEY + "hendrikpassword"), Role: "user", CustomerId: sql.NullInt32{Int32: 1, Valid: true}},
			expected: nil,
		},
		{
			name:     "Save user failed primary key required",
			want:     domain.Users{Password: bcryptPassword(config.SECRET_KEY + "hendrikpassword"), Role: "user", CustomerId: sql.NullInt32{Int32: 1, Valid: true}},
			expected: errs.NewUnexpectedError("error insert data user"),
		},
		{
			name:     "Save user failed primary key duplicate",
			want:     domain.Users{Username: "hendrik", Password: bcryptPassword(config.SECRET_KEY + "hendrikpassword"), Role: "user", CustomerId: sql.NullInt32{Int32: 1, Valid: true}},
			expected: errs.NewUnexpectedError("error insert data user"),
		},
		{
			name:     "Save user failed field empty password",
			want:     domain.Users{Username: "rizal", Role: "user", CustomerId: sql.NullInt32{Int32: 1, Valid: true}},
			expected: errs.NewUnexpectedError("error insert data user"),
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result := repository.SaveUser(testTable.want)
			assert.Equal(t, testTable.expected, result)
		})
	}
}
