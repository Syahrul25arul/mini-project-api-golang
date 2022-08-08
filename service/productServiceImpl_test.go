package service

import (
	"mini-project/database"
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/repostiory"
	"reflect"
	"testing"
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
