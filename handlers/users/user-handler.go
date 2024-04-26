package userhandler

import (
	"github.com/gin-gonic/gin"
	"main.go/schemas"
	"main.go/utils"
	"net/http"
)

func (h *userHandler) CreateUser(ctx *gin.Context) {
	var request schemas.CreateUserRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	user, err := h.userService.CreateUser(request)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}

func (h *userHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	_, user, err := h.userService.GetUserByID(id)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var request schemas.UpdateUserRequest

	err := ctx.ShouldBind(&request)
	if err != nil {

		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	user, _, err := h.userService.GetUserByID(id)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	user, err = h.userService.UpdateUser(request, *user)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}

func (h *userHandler) DeleteUserByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	_, _, err := h.userService.GetUserByID(ID)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	err = h.userService.DeleteUserByID(ID)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", nil, nil)
}
