package handler

import (
	"gin-simple-chat/entity"
	"gin-simple-chat/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler UserHandler) RegisterUser(ctx *gin.Context) {
	var request entity.UserRequest

	successResponse := entity.SuccessResponse{Error: false}
	errorResponse := entity.ErrorResponse{Error: true}

	err := ctx.Bind(&request)

	if err != nil {
		errorResponse.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	helper.HashPassword(&request.Password)

	result, err := handler.UserUsecase.RegisterUser(request)

	if err != nil {
		errorResponse.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	successResponse.Message = "Register User Succeed"
	successResponse.Data = result

	ctx.JSON(http.StatusCreated, successResponse)
}
