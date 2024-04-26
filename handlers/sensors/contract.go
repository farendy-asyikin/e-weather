package sensorhandler

import (
	"github.com/gin-gonic/gin"
	"main.go/services/sensors"
)

type SensorHandler interface {
	CreateSensor(ctx *gin.Context)
	GetSensorByID(ctx *gin.Context)
	UpdateSensor(ctx *gin.Context)
	DeleteSensor(ctx *gin.Context)
	//SoftDeleteSensor(ctx *gin.Context)
}

type sensorHandler struct {
	sensorService sensorservice.SensorService
}

func NewSensorHandler(sensorService sensorservice.SensorService) *sensorHandler {
	return &sensorHandler{sensorService}
}
