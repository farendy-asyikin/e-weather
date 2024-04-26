package deviceservice

import (
	"main.go/models"
	devicerepository "main.go/repositories/devices"
	"main.go/schemas"
)

type DeviceService interface {
	CreateDevice(request schemas.CreateDeviceRequest) (*models.Device, error)
	UpdateDevice(request schemas.UpdateDeviceRequest, device models.Device) (*models.Device, error)
	DeleteDevice(ID string) error
	GetDeviceByID(ID string) (*models.Device, *schemas.DetailDeviceResponse, error)
}

type deviceService struct {
	deviceRepository devicerepository.DeviceRepository
}

func NewDeviceService(
	deviceRepository devicerepository.DeviceRepository,
) *deviceService {
	return &deviceService{
		deviceRepository: deviceRepository,
	}
}
