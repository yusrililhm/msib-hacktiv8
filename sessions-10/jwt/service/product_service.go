package service

import (
	"go-jwt/entity"
	"go-jwt/model"
)

type ProductService interface {
	CreateProductService(userId uint, product *model.ProductModel) (*entity.Product, error)
	UpdateProductService(productId uint, userId uint, product *model.ProductModel) (*entity.Product, error)
}
