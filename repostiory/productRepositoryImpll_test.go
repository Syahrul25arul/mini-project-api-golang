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

func TestProductRepositoryImpl_GetAllProduct(t *testing.T) {
	// setup repository
	db := database.GetClientDb()
	repo := NewProductRepository(db)

	// create data product dummy
	repo.SetupProductDummy()

	// setup testCase
	testCase := []struct {
		name      string
		want      int
		expected  []domain.Product
		expected2 *errs.AppErr
	}{
		{
			name: "get all product default",
			want: 0,
			expected: []domain.Product{{
				ProductId:          1,
				ProductName:        "mie goreng sedap",
				CategoryId:         1,
				Price:              2500,
				Stock:              25,
				ProductDescription: "ini mie goreng sedap",
			},
				{
					ProductId:          2,
					ProductName:        "mie goreng udang sedap",
					CategoryId:         1,
					Price:              3200,
					Stock:              21,
					ProductDescription: "ini mie goreng udang sedap",
				},
				{
					ProductId:          3,
					ProductName:        "indomie kari ayam",
					CategoryId:         1,
					Price:              3000,
					Stock:              50,
					ProductDescription: "ini indomie kari ayam",
				},
				{
					ProductId:          4,
					ProductName:        "indomie goreng",
					CategoryId:         1,
					Price:              2700,
					Stock:              67,
					ProductDescription: "ini indomie goreng",
				},
				{
					ProductId:          5,
					ProductName:        "mie selera pedas",
					CategoryId:         1,
					Price:              2200,
					Stock:              31,
					ProductDescription: "ini mie selera pedas",
				},
			},
			expected2: nil,
		},
		{
			name: "get all product page 1",
			want: 1,
			expected: []domain.Product{{
				ProductId:          1,
				ProductName:        "mie goreng sedap",
				CategoryId:         1,
				Price:              2500,
				Stock:              25,
				ProductDescription: "ini mie goreng sedap",
			},
				{
					ProductId:          2,
					ProductName:        "mie goreng udang sedap",
					CategoryId:         1,
					Price:              3200,
					Stock:              21,
					ProductDescription: "ini mie goreng udang sedap",
				},
				{
					ProductId:          3,
					ProductName:        "indomie kari ayam",
					CategoryId:         1,
					Price:              3000,
					Stock:              50,
					ProductDescription: "ini indomie kari ayam",
				},
				{
					ProductId:          4,
					ProductName:        "indomie goreng",
					CategoryId:         1,
					Price:              2700,
					Stock:              67,
					ProductDescription: "ini indomie goreng",
				},
				{
					ProductId:          5,
					ProductName:        "mie selera pedas",
					CategoryId:         1,
					Price:              2200,
					Stock:              31,
					ProductDescription: "ini mie selera pedas",
				},
			},
			expected2: nil,
		},
		{
			name: "get all product page 2",
			want: 2,
			expected: []domain.Product{{
				ProductId:          6,
				ProductName:        "teh pucuk",
				CategoryId:         2,
				Price:              4000,
				Stock:              30,
				ProductDescription: "ini teh pucuk",
			},
				{
					ProductId:          7,
					ProductName:        "golda",
					CategoryId:         2,
					Price:              4000,
					Stock:              24,
					ProductDescription: "ini golda",
				},
				{
					ProductId:          8,
					ProductName:        "aqua",
					CategoryId:         2,
					Price:              5000,
					Stock:              26,
					ProductDescription: "ini aqua",
				},
				{
					ProductId:          9,
					ProductName:        "coca cola",
					CategoryId:         2,
					Price:              6000,
					Stock:              76,
					ProductDescription: "ini coca cola",
				},
				{
					ProductId:          10,
					ProductName:        "green tea",
					CategoryId:         2,
					Price:              8000,
					Stock:              38,
					ProductDescription: "ini green tea",
				},
			},
			expected2: nil,
		},
		{
			name: "get all product page 3",
			want: 3,
			expected: []domain.Product{{
				ProductId:          11,
				ProductName:        "rokok sampoerna",
				CategoryId:         3,
				Price:              28000,
				Stock:              38,
				ProductDescription: "ini rokok sampoerna",
			},
				{
					ProductId:          12,
					ProductName:        "rokok surya pro",
					CategoryId:         3,
					Price:              18000,
					Stock:              27,
					ProductDescription: "ini surya pro",
				},
				{
					ProductId:          13,
					ProductName:        "rokok Malboro",
					CategoryId:         3,
					Price:              34000,
					Stock:              38,
					ProductDescription: "ini rokok Malboro",
				},
				{
					ProductId:          14,
					ProductName:        "Gula 1kg",
					CategoryId:         3,
					Price:              14000,
					Stock:              18,
					ProductDescription: "ini Gula 1 kg",
				},
				{
					ProductId:          15,
					ProductName:        "Garam",
					CategoryId:         3,
					Price:              4000,
					Stock:              58,
					ProductDescription: "ini Garam",
				},
			},
			expected2: nil,
		},
		{
			name:      "get all product page 4",
			want:      4,
			expected:  []domain.Product{},
			expected2: nil,
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			products, err := repo.GetAllProduct(testTable.want)
			assert.Equal(t, testTable.expected, products)
			assert.Equal(t, testTable.expected2, err)
		})
	}
}

func TestProductRepositoryImpl_GetProductById(t *testing.T) {
	// setup repository
	db := database.GetClientDb()
	repo := NewProductRepository(db)

	// create data product dummy
	repo.SetupProductDummy()

	// setup testCase
	testCase := []struct {
		name      string
		want      string
		expected1 *domain.Product
		expected2 *errs.AppErr
	}{
		{
			name:      "get product by id = 1 success",
			want:      "1",
			expected1: &domain.Product{ProductId: 1, ProductName: "mie goreng sedap", CategoryId: 1, Price: 2500, Stock: 25, ProductDescription: "ini mie goreng sedap"},
			expected2: nil,
		},
		{
			name:      "get product by id = 5 success",
			want:      "6",
			expected1: &domain.Product{ProductId: 6, ProductName: "teh pucuk", CategoryId: 2, Price: 4000, Stock: 30, ProductDescription: "ini teh pucuk"},
			expected2: nil,
		},
		{
			name:      "get product by id not found",
			want:      "20",
			expected1: nil,
			expected2: errs.NewNotFoundError("product not found"),
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result, err := repo.GetProductById(testTable.want)
			assert.Equal(t, testTable.expected1, result)
			assert.Equal(t, testTable.expected2, err)
		})
	}
}

func TestProductRepositoryImpl_DeleteProduct(t *testing.T) {
	// setup repository
	db := database.GetClientDb()
	repo := NewProductRepository(db)

	// create data product dummy
	repo.SetupProductDummy()

	// setup testCase
	testCase := []struct {
		name     string
		want     string
		expected *errs.AppErr
	}{
		{
			name:     "delete product success",
			want:     "1",
			expected: nil,
		},
		{
			name:     "delete product not found",
			want:     "25",
			expected: errs.NewNotFoundError("product not found"),
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			err := repo.DeleteProduct(testTable.want)
			assert.Equal(t, testTable.expected, err)
		})
	}
}
