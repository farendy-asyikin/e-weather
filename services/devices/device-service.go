package deviceservice

import (
	"main.go/models"
	"main.go/schemas"
)

func (s *deviceService) CreateDevice(request schemas.CreateDeviceRequest) (*models.Device, error) {

	sensor := models.Device{
		Name: request.Name,
		Type: request.Type,
	}

	result, err := s.deviceRepository.CreateDevice(sensor)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *deviceService) GetDeviceByID(ID string) (*models.Device, *schemas.DetailDeviceResponse, error) {
	sensor, err := s.deviceRepository.GetDeviceByID(ID)
	if err != nil {
		return nil, nil, err
	}

	response := schemas.DetailDeviceResponse{
		Name:     sensor.Name,
		Type:     sensor.Type,
		IsActive: sensor.IsActive,
	}

	return sensor, &response, nil
}

func (s *deviceService) UpdateDevice(request schemas.UpdateDeviceRequest, device models.Device) (*models.Device, error) {

	result, err := s.deviceRepository.UpdateDevice(device)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *deviceService) DeleteDevice(ID string) error {
	err := s.deviceRepository.DeleteDevice(ID)
	if err != nil {
		return err
	}

	return nil
}
