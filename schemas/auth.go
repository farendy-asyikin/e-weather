package schemas

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}
