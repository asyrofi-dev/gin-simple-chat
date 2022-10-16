package usecase

import (
	"gin-simple-chat/entity"
	repository "gin-simple-chat/repository/user"
)

type IUserUsecase interface {
	RegisterUser(request entity.UserRequest) (entity.UserResponse, error)
}

type UserUsecase struct {
	UserRepo repository.IUserRepository
}

func NewUserUsecase(userRepo repository.IUserRepository) IUserUsecase {
	return UserUsecase{userRepo}
}
