package schemas

type CreateSensorRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type UpdateSensorRequest struct {
	Name *string `json:"name" binding:"required"`
	Type *string `json:"type" binding:"required"`
}

type DetailSensorResponse struct {
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required"`
	IsActive bool   `json:"is_active"`
}
