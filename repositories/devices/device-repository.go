package devicerepository

import (
	"errors"
	"gorm.io/gorm"
	"main.go/models"
)

func (r *deviceRepository) CreateDevice(sensor models.Device) (*models.Device, error) {

	err := r.db.Create(&sensor).Error
	if err != nil {
		return nil, err
	}

	return &sensor, nil
}

func (r *deviceRepository) GetDeviceByID(ID string) (*models.Device, error) {
	var device models.Device
	err := r.db.First(&device, ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("device not found")
		}

		return nil, err
	}

	return &device, nil
}

func (r *deviceRepository) UpdateDevice(device models.Device) (*models.Device, error) {
	err := r.db.Model(&device).Updates(map[string]any{
		"name":      device.Name,
		"type":      device.Type,
		"is_active": device.IsActive,
	}).Error
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *deviceRepository) DeleteDevice(ID string) error {
	var device models.Device
	err := r.db.Model(&device).Where("id = ?", ID).Updates(map[string]interface{}{"is_active": false}).Error
	if err != nil {
		return err
	}

	return nil
}
