package schemas

type Meta struct {
	Page       int64 `json:"page"`
	PerPage    int64 `json:"per_page"`
	TotalPages int64 `json:"total_pages"`
	TotalRows  int64 `json:"total_rows"`
}

type PaginationRequest struct {
	Page  int64 `form:"page" binding:"gt=0"`
	Limit int64 `form:"limit" binding:"gt=0"`
}
