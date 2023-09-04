package service

import (
	"go-jwt/entity"
	"go-jwt/model"
	"go-jwt/repository"
	"time"

	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
	}
}

// CreateProductService implements ProductService.
func (service *ProductServiceImpl) CreateProductService(userId uint, product *model.ProductModel) (*entity.Product, error) {
	products := &entity.Product{
		Title:       product.Title,
		Description: product.Description,
		UserId:      userId,
	}

	data, err := service.ProductRepository.CreateProductRepository(products)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateProductService implements ProductService.
func (service *ProductServiceImpl) UpdateProductService(productId uint, userId uint, product *model.ProductModel) (*entity.Product, error) {
	products := &entity.Product{
		Title:       product.Title,
		Description: product.Description,
		UserId:      userId,
		Model: gorm.Model{
			ID:        productId,
			UpdatedAt: time.Now(),
		},
	}

	data, err := service.ProductRepository.UpdateProductRepository(productId, products)

	if err != nil {
		return nil, err
	}

	return data, nil
}
