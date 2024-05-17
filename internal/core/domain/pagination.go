package domain

type Pagination struct {
	Page      int   `json:"page"`
	TotalData int64 `json:"total_data"`
}
