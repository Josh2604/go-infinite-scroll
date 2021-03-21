package getposts

type Paginator struct {
	PageNo int `json:"page_no" binding:"required"`
	Limit  int `json:"limit" binding:"required"`
}
