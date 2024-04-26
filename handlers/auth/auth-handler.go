package authhandler

import (
	"github.com/gin-gonic/gin"
	"main.go/schemas"
	"main.go/utils"
	"net/http"
)

func (h *authHandler) Login(ctx *gin.Context) {
	var input schemas.LoginRequest
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	response, err := h.authService.Login(input.Email, input.Password)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", response, nil)
}
