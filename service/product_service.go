package service

import (
	"challenge-3/entity"
	"challenge-3/repository"
	"errors"
	"fmt"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id int) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (service ProductService) GetAllProduct() (*[]entity.Product, error) {
	product := service.Repository.FindAll()
	fmt.Println(product)
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}
