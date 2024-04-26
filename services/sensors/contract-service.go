package sensorservice

import (
	"main.go/models"
	"main.go/schemas"
)

func (s *sensorService) CreateSensor(request schemas.CreateSensorRequest) (*models.Sensor, error) {

	sensor := models.Sensor{
		Name: request.Name,
		Type: request.Type,
	}

	result, err := s.sensorRepository.CreateSensor(sensor)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *sensorService) GetSensorByID(ID string) (*models.Sensor, *schemas.DetailSensorResponse, error) {
	sensor, err := s.sensorRepository.GetSensorByID(ID)
	if err != nil {
		return nil, nil, err
	}

	response := schemas.DetailSensorResponse{
		Name:     sensor.Name,
		Type:     sensor.Type,
		IsActive: sensor.IsActive,
	}

	return sensor, &response, nil
}

func (s *sensorService) UpdateSensor(request schemas.UpdateSensorRequest, sensor models.Sensor) (*models.Sensor, error) {

	result, err := s.sensorRepository.UpdateSensor(sensor)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *sensorService) DeleteSensor(ID string) error {
	err := s.sensorRepository.DeleteSensor(ID)
	if err != nil {
		return err
	}

	return nil
}
