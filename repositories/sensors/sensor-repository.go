package sensorrepository

import (
	"errors"
	"gorm.io/gorm"
	"main.go/models"
)

func (r *sensorRepository) CreateSensor(sensor models.Sensor) (*models.Sensor, error) {

	err := r.db.Create(&sensor).Error
	if err != nil {
		return nil, err
	}

	return &sensor, nil
}

func (r *sensorRepository) GetSensorByID(ID string) (*models.Sensor, error) {
	var sensor models.Sensor
	err := r.db.First(&sensor, ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("sensor not found")
		}

		return nil, err
	}

	return &sensor, nil
}

func (r *sensorRepository) UpdateSensor(sensor models.Sensor) (*models.Sensor, error) {
	err := r.db.Model(&sensor).Updates(map[string]any{
		"name":      sensor.Name,
		"type":      sensor.Type,
		"is_active": sensor.IsActive,
	}).Error
	if err != nil {
		return nil, err
	}

	return &sensor, nil
}

func (r *sensorRepository) DeleteSensor(ID string) error {
	var sensor models.Sensor
	err := r.db.Model(&sensor).Where("id = ?", ID).Updates(map[string]interface{}{"is_active": false}).Error
	if err != nil {
		return err
	}

	return nil
}
