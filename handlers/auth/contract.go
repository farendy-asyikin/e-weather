package authhandler

import (
	"github.com/gin-gonic/gin"
	authservice "main.go/services/auth"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
}

type authHandler struct {
	authService authservice.AuthService
}

func NewAuthHandler(authService authservice.AuthService) *authHandler {
	return &authHandler{authService}
}
