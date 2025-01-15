package requests

type PaginateRequest struct {
	Limit int `form:"limit,default=10" binding:"omitempty"`
	Page  int `form:"page,default=1" binding:"omitempty,min=1"`
}

type FilterRequest struct {
	Search  string `form:"search" binding:"omitempty,ascii"`
	OrderBy string `form:"order_by,default=created_at"`
	Sort    string `form:"sort,default=desc"`
}
