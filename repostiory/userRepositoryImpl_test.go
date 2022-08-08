package repostiory

import (
	"database/sql"
	"mini-project/database"
	"mini-project/domain"
	"mini-project/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestUserRepositoryImpl_SaveUser(t *testing.T) {
// 	// prepare database and repository
// 	db := database.GetClientDb()
// 	repository := NewUserRepository(db)

// 	customerRepo := NewCustomerRepository(db)
// 	newCustomer := domain.Customer{Name: "jamal", DateOfBirth: "1995-02-25", ZipCode: "13421", Status: "active"}
// 	customerRepo.Register(newCustomer)

// 	testCase := []struct {
// 		name     string
// 		want     domain.Users
// 		expected *errs.AppErr
// 	}{
// 		{
// 			name:     "Save user success",
// 			want:     domain.Users{Username: "hendrik", Password: bcryptPassword(config.SECRET_KEY + "hendrikpassword"), Role: "user", CustomerId: sql.NullInt32{Int32: 1, Valid: true}},
// 			expected: nil,
// 		},
// 		{
// 			name:     "Save user failed primary key required",
// 			want:     domain.Users{Password: bcryptPassword(config.SECRET_KEY + "hendrikpassword"), Role: "user", CustomerId: sql.NullInt32{Int32: 1, Valid: true}},
// 			expected: errs.NewUnexpectedError("error insert data user"),
// 		},
// 		{
// 			name:     "Save user failed primary key duplicate",
// 			want:     domain.Users{Username: "hendrik", Password: bcryptPassword(config.SECRET_KEY + "hendrikpassword"), Role: "user", CustomerId: sql.NullInt32{Int32: 1, Valid: true}},
// 			expected: errs.NewUnexpectedError("error insert data user"),
// 		},
// 		{
// 			name:     "Save user failed field empty password",
// 			want:     domain.Users{Username: "rizal", Role: "user", CustomerId: sql.NullInt32{Int32: 1, Valid: true}},
// 			expected: errs.NewUnexpectedError("error insert data user"),
// 		},
// 	}

// 	for _, testTable := range testCase {
// 		t.Run(testTable.name, func(t *testing.T) {
// 			result := repository.SaveUser(testTable.want)
// 			assert.Equal(t, testTable.expected, result)
// 		})
// 	}
// }

func TestGetAllUser(t *testing.T) {
	db := database.GetClientDb()
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
			expected1: &domain.Users{Username: "jamal12", Password: "$2a$08$rAUSlMpHl9OXIjVF9qIuBejkR5UfmpASCTxeDRradYcOhtdktpYPq", Role: "user", CustomerId: sql.NullInt32{Int32: 1, Valid: true}},
			expected2: nil,
		},
		{
			name:      "find user failed not found",
			want:      "test",
			expected1: nil,
			expected2: errs.NewNotFoundError("error get data user by username not found"),
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result, err := repo.FindByUsername(testTable.want)
			assert.Equal(t, result, testTable.expected1)
			assert.Equal(t, err, testTable.expected2)
		})
	}
}
