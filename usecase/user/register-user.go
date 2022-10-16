package usecase

import (
	"fmt"
	"gin-simple-chat/entity"
	"log"

	"github.com/jinzhu/copier"
)

func (usecase UserUsecase) RegisterUser(request entity.UserRequest) (entity.UserResponse, error) {
	var user entity.User
	var response entity.UserResponse

	copier.Copy(&user, &request)

	result, err := usecase.UserRepo.CreateUser(user)

	if err != nil {
		errMsg := fmt.Sprint("[UserUsecase] [RegisterUser] :", err)

		log.Println(errMsg)

		return entity.UserResponse{}, err
	}

	copier.Copy(&response, &result)

	return response, nil
}
