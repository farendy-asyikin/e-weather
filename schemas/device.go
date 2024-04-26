package schemas

type CreateDeviceRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type UpdateDeviceRequest struct {
	Name *string `json:"name" binding:"required"`
	Type *string `json:"type" binding:"required"`
}

type DetailDeviceResponse struct {
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required"`
	IsActive bool   `json:"is_active"`
}
