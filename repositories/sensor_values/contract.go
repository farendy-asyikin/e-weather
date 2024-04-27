package sensorvaluerepository

import (
	"gorm.io/gorm"
	"main.go/models"
	"main.go/schemas"
)

type SensorValueRepository interface {
	CreateValue(sensorValue models.SensorValue) (*models.SensorValue, error)
	GetValueByID(ID string) (*models.SensorValue, error)
	ListValuePagination(request schemas.ListValueRequest) (courier []*models.SensorValue, meta *schemas.Meta, err error)
	BulkDeleteValue(IDs []uint64) error
}

type sensorValueRepository struct {
	db *gorm.DB
}

func NewSensorValueRepository(db *gorm.DB) *sensorValueRepository {
	return &sensorValueRepository{db}
}
