package devicerepository

import (
	"gorm.io/gorm"
	"main.go/models"
)

type DeviceRepository interface {
	CreateDevice(device models.Device) (*models.Device, error)
	GetDeviceByID(ID string) (*models.Device, error)
	UpdateDevice(device models.Device) (*models.Device, error)
	DeleteDevice(ID string) error
}

type deviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) *deviceRepository {
	return &deviceRepository{db}
}
