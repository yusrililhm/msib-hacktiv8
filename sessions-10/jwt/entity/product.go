package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string `json:"title" valid:"required~ Title of your product is required"`
	Description string `json:"description" valid:"required~ Description of your product is required"`
	UserId      uint   `json:"userId"`
	User        *User
}

func (p *Product) BeforeCreate(db *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return err
	}

	return nil
}

func (p *Product) BeforeUpdate(db *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return err
	}

	return nil
}
