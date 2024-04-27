package sensorvaluehandler

import (
	"github.com/gin-gonic/gin"
	sensorvalueservice "main.go/services/sensor_values"
)

type SensorValueHandler interface {
	CreateValue(ctx *gin.Context)
	GetValueByID(ctx *gin.Context)
	ListValuePagination(ctx *gin.Context)
	BulkDeleteValue(ctx *gin.Context)
}

type sensorValueHandler struct {
	sensorValueService sensorvalueservice.SensorValueService
}

func NewSensorValueHandler(sensorValueService sensorvalueservice.SensorValueService) *sensorValueHandler {
	return &sensorValueHandler{sensorValueService}
}
