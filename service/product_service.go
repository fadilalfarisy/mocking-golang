package service

import (
	"challenge-3/entity"
	"challenge-3/repository"
	"errors"
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

func (service ProductService) GetAllProduct(status string) (*[]entity.Product, error) {
	product := service.Repository.FindAll(status)

	if len(*product) == 0 {
		return nil, errors.New("product not found")
	}

	return product, nil
}
