package repository

import "challenge-3/entity"

type ProductRepository interface {
	FindById(id int) *entity.Product
	FindAll(status string) *[]entity.Product
}
