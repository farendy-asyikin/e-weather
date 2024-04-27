package sensorvaluehandler

import (
	"github.com/gin-gonic/gin"
	"main.go/schemas"
	"main.go/utils"
	"net/http"
)

func (h *sensorValueHandler) CreateValue(ctx *gin.Context) {
	var request schemas.CreateValueRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	user, err := h.sensorValueService.CreateValue(request)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}
func (h *sensorValueHandler) GetValueByID(ctx *gin.Context) {
	id := ctx.Param("id")

	_, user, err := h.sensorValueService.GetValueByID(id)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", user, nil)
}
func (h *sensorValueHandler) ListValuePagination(ctx *gin.Context) {
	var request schemas.ListValueRequest
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	if request.OrderBy != nil && request.Order != nil {
		// validate order is asc or desc
		if *request.Order != "asc" && *request.Order != "desc" {
			utils.ApiResponse(ctx, http.StatusBadRequest, "order must be asc or desc", nil, nil)
			return
		}
	}

	couriers, meta, err := h.sensorValueService.ListValuePagination(request)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", couriers, meta)
}

func (h *sensorValueHandler) BulkDeleteValue(ctx *gin.Context) {
	var request schemas.BulkDeleteRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		error := utils.FormatValidationError(err)
		utils.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	err = h.sensorValueService.BulkDeleteValue(request.IDs)
	if err != nil {
		utils.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utils.ApiResponse(ctx, http.StatusOK, "success", nil, nil)
}
