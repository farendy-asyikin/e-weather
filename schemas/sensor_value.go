package schemas

type CreateValueRequest struct {
	SensorID uint64  `json:"sensor_id" binding:"required"`
	Value    float64 `json:"value" binding:"required"`
}

type DetailValueResponse struct {
	SensorID   uint64  `json:"sensor_id"`
	SensorName string  `json:"sensor_name"`
	Value      float64 `json:"value"`
}

type ListValueRequest struct {
	ListPaginationRequest
}

type ListValueResponse struct {
	SensorID   uint64  `json:"sensor_id"`
	SensorName string  `json:"sensor_name"`
	Value      float64 `json:"value"`
}
