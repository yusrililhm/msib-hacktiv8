package repository

import (
	"go-jwt/entity"
)

type ProductRepository interface {
	CreateProductRepository(product *entity.Product) (*entity.Product, error)
	UpdateProductRepository(productId uint, product *entity.Product) (*entity.Product, error)
}
