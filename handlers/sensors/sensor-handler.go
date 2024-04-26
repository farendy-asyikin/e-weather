package sensorhandler

import (
	"github.com/gin-gonic/gin"
	"main.go/schemas"
	"main.go/utils"
	"net/http"
)

//SoftDeleteSensor(ctx *gin.Context)

func (h *sensorHandler) CreateSensor(ctx *gin.Context) {
	var request schemas.CreateSensorRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	user, err := h.sensorService.CreateSensor(request)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}

func (h *sensorHandler) GetSensorByID(ctx *gin.Context) {
	id := ctx.Param("id")

	_, user, err := h.sensorService.GetSensorByID(id)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}

func (h *sensorHandler) UpdateSensor(ctx *gin.Context) {
	id := ctx.Param("id")
	var request schemas.UpdateSensorRequest

	err := ctx.ShouldBind(&request)
	if err != nil {

		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	user, _, err := h.sensorService.GetSensorByID(id)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	user, err = h.sensorService.UpdateSensor(request, *user)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}

func (h *sensorHandler) DeleteSensor(ctx *gin.Context) {
	ID := ctx.Param("id")

	_, _, err := h.sensorService.GetSensorByID(ID)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	err = h.sensorService.DeleteSensor(ID)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", nil, nil)
}
