package repository

import (
	"go-jwt/entity"
)

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	Login(email string) (*entity.User, error)
}
