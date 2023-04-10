package service

import (
	"challenge-3/entity"
	"challenge-3/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {

	productRepository.Mock.On("FindById", 1).Return(nil)

	product, err := productService.GetOneProduct(1)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := entity.Product{
		Id:          2,
		Title:       "Kaca mata",
		Description: "Lensa Photocromic",
	}

	productRepository.Mock.On("FindById", 2).Return(product)

	result, err := productService.GetOneProduct(2)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, product.Id, result.Id, "result has to be 2")
	assert.Equal(t, product.Title, result.Title, "result has to be 'Kaca mata'")
	assert.Equal(t, product.Description, result.Description, "result has to be 'Lensa Photocromic'")
	assert.Equal(t, &product, result, "result has to be a product with id 2")
}

func TestProductServiceGetAllProduct(t *testing.T) {
	products := []entity.Product{
		{
			Id:          1,
			Title:       "Kaca mata",
			Description: "Lensa Photocromic",
		},
		{
			Id:          2,
			Title:       "Mouse",
			Description: "Lampu RGB",
		},
	}

	productRepository.Mock.On("FindAll").Return(products)

	result, err := productService.GetAllProduct()

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, (*result)[0], entity.Product{
		Id:          1,
		Title:       "Kaca mata",
		Description: "Lensa Photocromic",
	}, "result has id 1, title kaca mata, and description lensa photocromic")

	assert.Equal(t, (*result)[1], entity.Product{
		Id:          2,
		Title:       "Mouse",
		Description: "Lampu RGB",
	}, "result has id 2, title mouse, and description lampu rgb")
}

func TestProductServiceGetAllProductNotFound(t *testing.T) {
	product := []entity.Product{}
	productRepository.Mock.On("FindAll").Return(product)

	result, err := productService.GetAllProduct()

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}
