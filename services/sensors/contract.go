package sensorservice

import (
	"main.go/models"
	sensorrepository "main.go/repositories/sensors"
	"main.go/schemas"
)

type SensorService interface {
	CreateSensor(request schemas.CreateSensorRequest) (*models.Sensor, error)
	UpdateSensor(request schemas.UpdateSensorRequest, sensor models.Sensor) (*models.Sensor, error)
	DeleteSensor(ID string) error
	GetSensorByID(ID string) (*models.Sensor, *schemas.DetailSensorResponse, error)
}

type sensorService struct {
	sensorRepository sensorrepository.SensorRepository
}

func NewSensorService(
	sensorRepository sensorrepository.SensorRepository,
) *sensorService {
	return &sensorService{
		sensorRepository: sensorRepository,
	}
}
