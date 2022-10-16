package router

import (
	"gin-simple-chat/config"
	handler "gin-simple-chat/handler/user"
	repository "gin-simple-chat/repository/user"
	usecase "gin-simple-chat/usecase/user"

	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.Engine) {
	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	g.POST("/users", userHandler.RegisterUser)
}
