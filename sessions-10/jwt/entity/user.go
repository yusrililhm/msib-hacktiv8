package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represent the model for an user
type User struct {
	gorm.Model
	FullName string    `json:"fullName" gorm:"not null;" valid:"required~ Your full name is required"`
	Email    string    `json:"email" gorm:"not null; uniqueIndex" valid:"required~ Your email is required, email~ invalid email format"`
	Password string    `json:"password" gorm:"not null;" valid:"required~ Your password is required~  Password has to have a minimun length 6 characters"`
	Products []Product `json:"products" gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return err
	}

	return nil
}
