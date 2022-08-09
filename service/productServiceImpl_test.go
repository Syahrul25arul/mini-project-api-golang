package service

import (
	"mini-project/database"
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/repostiory"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductServiceImpl_SaveProductService(t *testing.T) {
	db := database.GetClientDb()
	service := NewProductService(repostiory.NewProductRepository(db))
	type args struct {
		product domain.Product
	}
	tests := []struct {
		name string
		s    ProductServiceImpl
		args args
		want *errs.AppErr
	}{
		{
			name: "product valid",
			s:    service,
			args: args{product: domain.Product{ProductName: "teh pucuk", CategoryId: 2, Price: 10000, Stock: 20, ProductDescription: "ini teh pucuk"}},
			want: nil,
		},
		{
			name: "product valid description null",
			s:    service,
			args: args{product: domain.Product{ProductName: "teh pucuk", CategoryId: 2, Price: 10000, Stock: 20, ProductDescription: ""}},
			want: nil,
		},
		{
			name: "product failed duplicate primary key",
			s:    service,
			args: args{product: domain.Product{ProductId: 1, ProductName: "teh pucuk", CategoryId: 2, Price: 10000, Stock: 20, ProductDescription: "ini teh pucuk"}},
			want: errs.NewUnexpectedError("error insert data product"),
		},
		{
			name: "product invalid categoryId not less than 1",
			s:    service,
			args: args{product: domain.Product{ProductName: "teh pucuk", CategoryId: 0, Price: 10000, Stock: 20, ProductDescription: "ini teh pucuk"}},
			want: errs.NewValidationError("field CategoryId cannot less than 1"),
		},
		{
			name: "product invalid field price not less than 0",
			s:    service,
			args: args{product: domain.Product{ProductName: "teh pucuk", CategoryId: 2, Price: -1, Stock: 20, ProductDescription: "ini teh pucuk"}},
			want: errs.NewValidationError("field Price cannot less than 0"),
		},
		{
			name: "product invalid field stock not less than 0",
			s:    service,
			args: args{product: domain.Product{ProductName: "teh pucuk", CategoryId: 2, Price: 4000, Stock: -1, ProductDescription: "ini teh pucuk"}},
			want: errs.NewValidationError("field Stock cannot less than 0"),
		},
		{
			name: "product invalid field productName cannot be empty",
			s:    service,
			args: args{product: domain.Product{CategoryId: 2, Price: 4000, Stock: 20, ProductDescription: "ini teh pucuk"}},
			want: errs.NewValidationError("field ProductName cannot be empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SaveProductService(tt.args.product); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductServiceImpl.SaveProductService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductServiceImpl_isValid(t *testing.T) {
	db := database.GetClientDb()
	service := NewProductService(repostiory.NewProductRepository(db))
	type args struct {
		product domain.Product
	}
	tests := []struct {
		name string
		s    ProductServiceImpl
		args args
		want *errs.AppErr
	}{
		{
			name: "product valid",
			s:    service,
			args: args{product: domain.Product{ProductName: "teh pucuk", CategoryId: 2, Price: 10000, Stock: 20, ProductDescription: "ini teh pucuk"}},
			want: nil,
		},
		{
			name: "product valid description null",
			s:    service,
			args: args{product: domain.Product{ProductName: "teh pucuk", CategoryId: 2, Price: 10000, Stock: 20, ProductDescription: ""}},
			want: nil,
		},
		{
			name: "product invalid categoryId not less than 1",
			s:    service,
			args: args{product: domain.Product{ProductName: "teh pucuk", CategoryId: 0, Price: 10000, Stock: 20, ProductDescription: "ini teh pucuk"}},
			want: errs.NewValidationError("field CategoryId cannot less than 1"),
		},
		{
			name: "product invalid field price not less than 0",
			s:    service,
			args: args{product: domain.Product{ProductName: "teh pucuk", CategoryId: 2, Price: -1, Stock: 20, ProductDescription: "ini teh pucuk"}},
			want: errs.NewValidationError("field Price cannot less than 0"),
		},
		{
			name: "product invalid field stock not less than 0",
			s:    service,
			args: args{product: domain.Product{ProductName: "teh pucuk", CategoryId: 2, Price: 4000, Stock: -1, ProductDescription: "ini teh pucuk"}},
			want: errs.NewValidationError("field Stock cannot less than 0"),
		},
		{
			name: "product invalid field productName cannot be empty",
			s:    service,
			args: args{product: domain.Product{CategoryId: 2, Price: 4000, Stock: 20, ProductDescription: "ini teh pucuk"}},
			want: errs.NewValidationError("field ProductName cannot be empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.isValid(tt.args.product); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductServiceImpl.isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductServiceImpl_GetAllProduct(t *testing.T) {
	// setup service
	db := database.GetClientDb()
	repo := repostiory.NewProductRepository(db)
	service := NewProductService(repo)

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
			products, err := service.GetAllProductService(testTable.want)
			assert.Equal(t, testTable.expected, products)
			assert.Equal(t, testTable.expected2, err)
		})
	}
}

func TestProductServiceImpl_GetProductByIdService(t *testing.T) {
	// setup service
	db := database.GetClientDb()
	repo := repostiory.NewProductRepository(db)
	service := NewProductService(repo)

	// create data product dummy
	repo.SetupProductDummy()

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
			result, err := service.GetProductByIdService(testTable.want)
			assert.Equal(t, testTable.expected1, result)
			assert.Equal(t, testTable.expected2, err)
		})
	}
}

func TestProductServiceImpl_DeleteProductService(t *testing.T) {
	// setup service
	db := database.GetClientDb()
	repo := repostiory.NewProductRepository(db)
	service := NewProductService(repo)

	// create data product dummy
	repo.SetupProductDummy()

	// setup testCase
	testCase := []struct {
		name     string
		want     string
		expected *errs.AppErr
	}{
		{
			name:     "delete product service success",
			want:     "1",
			expected: nil,
		},
		{
			name:     "delete product service not found",
			want:     "25",
			expected: errs.NewNotFoundError("product not found"),
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			err := service.DeleteProductService(testTable.want)
			assert.Equal(t, testTable.expected, err)
		})
	}
}
