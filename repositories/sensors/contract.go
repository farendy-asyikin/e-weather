package sensorrepository

import (
	"gorm.io/gorm"
	"main.go/models"
)

type SensorRepository interface {
	CreateSensor(sensor models.Sensor) (*models.Sensor, error)
	GetSensorByID(ID string) (*models.Sensor, error)
	UpdateSensor(sensor models.Sensor) (*models.Sensor, error)
	DeleteSensor(ID string) error
}

type sensorRepository struct {
	db *gorm.DB
}

func NewSensorRepository(db *gorm.DB) *sensorRepository {
	return &sensorRepository{db}
}
