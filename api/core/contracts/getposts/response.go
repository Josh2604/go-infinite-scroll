package getposts

import "github.com/Josh2604/go-infinite-scroll/api/core/entities"

type Response struct {
	Total       int             `json:"total"`
	CurrentPage int             `json:"current_page"`
	PagesNo     int             `json:"pages_no"`
	HasMore     bool            `json:"has_more"`
	Items       []entities.Post `json:"items"`
}
