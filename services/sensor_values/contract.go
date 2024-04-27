package sensorvalueservice

import (
	"main.go/models"
	sensorvaluerepository "main.go/repositories/sensor_values"
	"main.go/schemas"
)

type SensorValueService interface {
	CreateValue(request schemas.CreateValueRequest) (*models.SensorValue, error)
	GetValueByID(ID string) (*models.SensorValue, *schemas.DetailValueResponse, error)
	ListValuePagination(request schemas.ListValueRequest) (response []*schemas.ListValueResponse, meta *schemas.Meta, err error)
	BulkDeleteValue(IDs []uint64) error
}

type sensorValueService struct {
	sensorValueRepository sensorvaluerepository.SensorValueRepository
}

func NewSensorValueService(
	sensorValueRepository sensorvaluerepository.SensorValueRepository,
) *sensorValueService {
	return &sensorValueService{
		sensorValueRepository: sensorValueRepository,
	}
}
