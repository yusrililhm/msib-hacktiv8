package service

import (
	"go-jwt/entity"
	"go-jwt/model"
	"go-jwt/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

// UserRegisterService implements UserService.
func (service *UserServiceImpl) UserRegisterService(user *model.UserRegister) (*model.UserModel, error) {

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	users := &entity.User{
		FullName: user.FullName,
		Email:    user.Email,
		Password: string(password),
	}

	_, err := service.UserRepository.Save(users)

	if err != nil {
		return nil, err
	}

	userResponse := &model.UserModel{
		Id:       users.ID,
		FullName: users.FullName,
		Email:    users.Email,
	}

	return userResponse, nil
}

// UserLoginService implements UserService.
func (service *UserServiceImpl) UserLoginService(user *model.UserLogin) (*model.UserModel, error) {
	data, err := service.UserRepository.Login(user.Email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(user.Password))

	if err != nil {
		return nil, err
	}

	return &model.UserModel{
		Id:       data.ID,
		FullName: data.FullName,
		Email:    data.Email,
	}, nil
}
