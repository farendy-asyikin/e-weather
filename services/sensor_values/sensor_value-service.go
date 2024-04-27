package sensorvalueservice

import (
	"main.go/models"
	"main.go/schemas"
)

func (s *sensorValueService) CreateValue(request schemas.CreateValueRequest) (*models.SensorValue, error) {
	value := models.SensorValue{
		SensorID: request.SensorID,
		Value:    request.Value,
	}

	result, err := s.sensorValueRepository.CreateValue(value)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *sensorValueService) GetValueByID(ID string) (*models.SensorValue, *schemas.DetailValueResponse, error) {
	value, err := s.sensorValueRepository.GetValueByID(ID)
	if err != nil {
		return nil, nil, err
	}

	response := schemas.DetailValueResponse{
		SensorID:   value.Sensor.ID,
		SensorName: value.Sensor.Name,
		Value:      value.Value,
	}

	return value, &response, nil
}
func (s *sensorValueService) ListValuePagination(request schemas.ListValueRequest) (response []*schemas.ListValueResponse, meta *schemas.Meta, err error) {
	values, meta, err := s.sensorValueRepository.ListValuePagination(request)
	if err != nil {
		return nil, nil, err
	}

	for _, value := range values {
		response = append(response, &schemas.ListValueResponse{
			SensorID:   value.Sensor.ID,
			SensorName: value.Sensor.Name,
			Value:      value.Value,
		})
	}

	return response, meta, nil
}

func (s *sensorValueService) BulkDeleteValue(IDs []uint64) error {
	err := s.sensorValueRepository.BulkDeleteValue(IDs)
	if err != nil {
		return err
	}

	return nil
}
