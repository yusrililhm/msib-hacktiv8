package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Brand     string    `json:"brand" gorm:"not null"`
	UserId    uint      `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ProductRepository struct {
	*gorm.DB
}

type ProductRepositoryImpl interface {
	CreateProduct(product Product) (Product, error)
}

func NewProductRepository(db *gorm.DB) ProductRepositoryImpl {
	return &ProductRepository{
		DB: db,
	}
}

func (p *Product) BeforeCreate(db *gorm.DB) (err error) {
	if len(p.Name) < 4 {
		err = errors.New("name too short")
	}
	return
}

func (r *ProductRepository) CreateProduct(product Product) (Product, error) {
	if err := r.Create(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}
