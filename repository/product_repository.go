package repository

import "golang-jwt-auth/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
