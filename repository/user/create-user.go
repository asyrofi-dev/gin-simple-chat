package repository

import (
	"fmt"
	"gin-simple-chat/entity"
	"log"
)

func (repo UserRepository) CreateUser(user entity.User) (entity.User, error) {
	query := repo.DB.Debug().Create(&user)

	if query.Error != nil {
		errMsg := fmt.Sprint("[UserRepository] [CreateUser] : ", query.Error)

		log.Println(errMsg)

		return entity.User{}, query.Error
	}

	return user, nil
}
