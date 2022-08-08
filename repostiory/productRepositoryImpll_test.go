package repostiory

import (
	"mini-project/database"
	"mini-project/domain"
	"mini-project/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductRepositoryImpl_SaveProduct(t *testing.T) {
	// setup repository
	db := database.GetClientDb()
	repo := NewProductRepository(db)

	// setup testCase
	testCase := []struct {
		name     string
		want     *domain.Product
		expected *errs.AppErr
	}{
		{
			name:     "save product success",
			want:     &domain.Product{ProductName: "Mie Goreng", CategoryId: 1, Price: 12000, ProductDescription: "ini mie goreng", Stock: 20},
			expected: nil,
		},
		{
			name:     "save product failed error duplicate key value",
			want:     &domain.Product{ProductId: 1, ProductName: "Mie Goreng", CategoryId: 1, Price: 12000, ProductDescription: "ini mie goreng", Stock: 20},
			expected: errs.NewUnexpectedError("error insert data product"),
		},
		{
			name:     "save product failed error field price not < 0",
			want:     &domain.Product{ProductName: "Teh pucuk", CategoryId: 2, ProductDescription: "ini teh pucuk", Price: -1, Stock: 15},
			expected: errs.NewUnexpectedError("error insert data product"),
		},
		{
			name:     "save product failed error field price stock not < 0",
			want:     &domain.Product{ProductName: "Teh pucuk", CategoryId: 2, ProductDescription: "ini teh pucuk", Price: 20, Stock: -1},
			expected: errs.NewUnexpectedError("error insert data product"),
		},
	}

	// run test case
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result := repo.SaveProduct(testTable.want)
			assert.Equal(t, testTable.expected, result)
		})
	}
}
