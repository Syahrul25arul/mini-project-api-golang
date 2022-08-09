package service

import (
	"fmt"
	"mini-project/domain"
	"mini-project/errs"
	"mini-project/logger"
	"mini-project/repostiory"
	"reflect"
)

type ProductServiceImpl struct {
	repo repostiory.ProductRepository
}

func NewProductService(repo repostiory.ProductRepository) ProductServiceImpl {
	return ProductServiceImpl{repo}
}

func (s ProductServiceImpl) SaveProductService(product domain.Product) *errs.AppErr {
	// cek apakah data yang dimasukkan sudah valid
	if err := s.isValid(product); err != nil {
		logger.Error("error invalid data customer " + err.Message)
		return err
	}
	return s.repo.SaveProduct(&product)

}

func (s ProductServiceImpl) GetAllProductService(page int) ([]domain.Product, *errs.AppErr) {
	return s.repo.GetAllProduct(page)
}

func (s ProductServiceImpl) GetProductByIdService(productId string) (*domain.Product, *errs.AppErr) {
	return s.repo.GetProductById(productId)
}

func (s ProductServiceImpl) isValid(product domain.Product) *errs.AppErr {
	// buat reflection untuk domain register
	ref := reflect.TypeOf(product)

	// cek field product
	for i := 0; i < ref.NumField(); i++ {

		if ref.Field(i).Name != "ProductId" && (ref.Field(i).Name == "Price" || ref.Field(i).Name == "Stock") && reflect.ValueOf(product).Field(i).Interface().(int64) < 0 {
			// cek apakah price atau stock kurang dari 0
			return errs.NewValidationError(fmt.Sprintf("field %s cannot less than 0", ref.Field(i).Name))
		} else if ref.Field(i).Name == "CategoryId" && reflect.ValueOf(product).Field(i).Interface().(int64) < 1 {
			// cek apakah value CategoryId kurang dari 1
			return errs.NewValidationError(fmt.Sprintf("field %s cannot less than 1", ref.Field(i).Name))
		} else if reflect.ValueOf(product).Field(i).Interface() == "" && ref.Field(i).Name != "ProductDescription" {
			// cek apakah value ProductName = empty string
			return errs.NewValidationError(fmt.Sprintf("field %s cannot be empty", ref.Field(i).Name))
		}
	}

	return nil
}
