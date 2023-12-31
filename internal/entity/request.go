package entity

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
