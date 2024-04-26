package userhandler

import (
	"github.com/gin-gonic/gin"
	userservice "main.go/services/user"
)

type UserHandler interface {
	CreateUser(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUserByID(ctx *gin.Context)
}

type userHandler struct {
	userService userservice.UserService
}

func NewUserHandler(userService userservice.UserService) *userHandler {
	return &userHandler{userService}
}
