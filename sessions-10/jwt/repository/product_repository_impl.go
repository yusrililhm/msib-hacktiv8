package repository

import (
	"go-jwt/entity"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	*gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

// CreateProductRepository implements ProductRepository.
func (repository *ProductRepositoryImpl) CreateProductRepository(product *entity.Product) (*entity.Product, error) {
	if err := repository.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

// UpdateProductRepository implements ProductRepository.
func (repository *ProductRepositoryImpl) UpdateProductRepository(productId uint, product *entity.Product) (*entity.Product, error) {
	if err := repository.Model(product).Where("id = ?", productId).Updates(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
