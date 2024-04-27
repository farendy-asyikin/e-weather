package sensorvaluerepository

import (
	"errors"
	"gorm.io/gorm"
	"main.go/models"
	"main.go/schemas"
	"math"
)

func (r *sensorValueRepository) CreateValue(sensorValue models.SensorValue) (*models.SensorValue, error) {

	err := r.db.Create(&sensorValue).Error
	if err != nil {
		return nil, err
	}

	return &sensorValue, nil
}
func (r *sensorValueRepository) GetValueByID(ID string) (*models.SensorValue, error) {
	var sensorValue models.SensorValue
	err := r.db.Preload("Sensor").First(&sensorValue, ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("value not found")
		}

		return nil, err
	}

	return &sensorValue, nil
}

func (r *sensorValueRepository) ListValuePagination(request schemas.ListValueRequest) (couriers []*models.SensorValue, meta *schemas.Meta, err error) {
	var (
		offset = (request.Page - 1) * request.Limit
		db     = r.db.Model(&couriers)
		count  int64
	)

	if request.Search != nil && *request.Search != "" {
		search := "%" + *request.Search + "%"
		db = db.Where("name LIKE ?", search)
	}

	if err := db.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	db.Offset(int(offset)).Limit(int(request.Limit))
	if request.OrderBy != nil && request.Order != nil {
		db = db.Order(*request.OrderBy + " " + *request.Order)
	} else {
		db = db.Order("created_at DESC")
	}

	if err := db.Find(&couriers).Error; err != nil {
		return nil, nil, err
	}

	meta = &schemas.Meta{
		Page:       request.Page,
		PerPage:    request.Limit,
		TotalPages: int64(math.Ceil(float64(count) / float64(request.Limit))),
		TotalRows:  count,
	}

	return couriers, meta, nil
}

func (r *sensorValueRepository) BulkDeleteValue(IDs []uint64) error {
	err := r.db.Delete(&models.SensorValue{}, IDs).Error
	if err != nil {
		return err
	}

	return nil
}
