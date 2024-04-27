package schemas

type BulkDeleteRequest struct {
	IDs []uint64 `json:"ids" binding:"required"`
}

type ListPaginationRequest struct {
	PaginationRequest
	Search  *string `form:"search"`
	OrderBy *string `form:"order_by"`
	Order   *string `form:"order"`
}
