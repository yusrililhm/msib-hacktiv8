package repository

import (
	"go-jwt/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) Save(user *entity.User) (*entity.User, error) {
	if err := repository.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Login implements UserRepository.
func (repository *UserRepositoryImpl) Login(email string) (*entity.User, error) {
	users := &entity.User{}

	if err := repository.Model(users).First(users, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return users, nil
}
