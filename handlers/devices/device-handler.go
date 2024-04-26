package devicehandler

import (
	"github.com/gin-gonic/gin"
	"main.go/schemas"
	"main.go/utils"
	"net/http"
)

//SoftDeleteDevice(ctx *gin.Context)

func (h *deviceHandler) CreateDevice(ctx *gin.Context) {
	var request schemas.CreateDeviceRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	user, err := h.deviceService.CreateDevice(request)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}

func (h *deviceHandler) GetDeviceByID(ctx *gin.Context) {
	id := ctx.Param("id")

	_, user, err := h.deviceService.GetDeviceByID(id)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}

func (h *deviceHandler) UpdateDevice(ctx *gin.Context) {
	id := ctx.Param("id")
	var request schemas.UpdateDeviceRequest

	err := ctx.ShouldBind(&request)
	if err != nil {

		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	user, _, err := h.deviceService.GetDeviceByID(id)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	user, err = h.deviceService.UpdateDevice(request, *user)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}

func (h *deviceHandler) DeleteDevice(ctx *gin.Context) {
	ID := ctx.Param("id")

	_, _, err := h.deviceService.GetDeviceByID(ID)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	err = h.deviceService.DeleteDevice(ID)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", nil, nil)
}
