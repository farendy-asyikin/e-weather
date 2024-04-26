package devicehandler

import (
	"github.com/gin-gonic/gin"
	deviceservice "main.go/services/devices"
)

type DeviceHandler interface {
	CreateDevice(ctx *gin.Context)
	GetDeviceByID(ctx *gin.Context)
	UpdateDevice(ctx *gin.Context)
	DeleteDevice(ctx *gin.Context)
	//SoftDeleteDevice(ctx *gin.Context)
}

type deviceHandler struct {
	deviceService deviceservice.DeviceService
}

func NewDeviceHandler(deviceService deviceservice.DeviceService) *deviceHandler {
	return &deviceHandler{deviceService}
}
