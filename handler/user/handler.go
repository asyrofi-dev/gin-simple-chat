package handler

import usecase "gin-simple-chat/usecase/user"

type UserHandler struct {
	UserUsecase usecase.IUserUsecase
}

func NewUserHandler(userUsecase usecase.IUserUsecase) UserHandler {
	return UserHandler{userUsecase}
}
