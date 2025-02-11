package models

type Pagination struct {
	Total       int `json:"total,omitempty"`
	Count       int `json:"count,omitempty"`
	PerPage     int `json:"per_page,omitempty"`
	CurrentPage int `json:"current_page,omitempty"`
	TotalPages  int `json:"total_pages,omitempty"`
}
