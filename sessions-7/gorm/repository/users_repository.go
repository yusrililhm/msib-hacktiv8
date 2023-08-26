package repository

import (
	"fmt"
	"go-gorm/model"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	*gorm.DB
}

type UserRepositoryImpl interface {
	CreateUser(user model.User) (model.User, error)
	GetUser(id int) (model.User, error)
	UserUpdate(id int, user model.User) (string, error)
	GetUserWithProduct(user model.User) (model.User, error)
}

func NewRepository(db *gorm.DB) UserRepositoryImpl {
	return &UserRepository{
		DB: db,
	}
}

// CreateUser implements UserRepositoryImpl.
func (repository *UserRepository) CreateUser(user model.User) (model.User, error) {

	users := model.User{
		Email: user.Email,
	}

	if err := repository.Create(&users).Error; err != nil {
		log.Fatal(err.Error())
		return users, err
	}

	return users, nil
}

// GetUser implements UserRepositoryImpl.
func (repository *UserRepository) GetUser(id int) (model.User, error) {
	user := model.User{}

	if err := repository.First(&user, "id = $1", id).Error; err != nil {
		return user, err
	}

	return user, nil
}

// UserUpdate implements UserRepositoryImpl.
func (repository *UserRepository) UserUpdate(id int, user model.User) (string, error) {
	if err := repository.Model(&model.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return fmt.Sprintf("Update fail user id %d not found", id), err
	}
	return fmt.Sprintf("Update user id %d success", id), nil
}

// GetUserWithProduct implements UserRepositoryImpl.
func (repository *UserRepository) GetUserWithProduct(user model.User) (model.User, error) {
	if err := repository.Preload("Products").Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
