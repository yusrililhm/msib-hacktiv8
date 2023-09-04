package service

import (
	"go-jwt/model"
)

type UserService interface {
	UserRegisterService(user *model.UserRegister) (*model.UserModel, error)
	UserLoginService(user *model.UserLogin) (*model.UserModel, error)
}
