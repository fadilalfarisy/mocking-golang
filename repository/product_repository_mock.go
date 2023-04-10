package repository

import (
	"challenge-3/entity"
	"log"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id int) *entity.Product {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	log.Println(arguments.Get(0).(entity.Product))
	product := arguments.Get(0).(entity.Product)
	return &product
}

func (repository *ProductRepositoryMock) FindAll() *[]entity.Product {
	arguments := repository.Mock.Called()

	product := arguments.Get(0).([]entity.Product)
	log.Println(product)
	return &product
}
