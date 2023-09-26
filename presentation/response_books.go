package presentation

import "time"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Category    []string   `json:"category"`
	Keyword     string     `json:"keyword"`
	Price       string     `json:"price"`
	Stock       int64      `json:"stock"`
	Publisher   string     `json:"publisher"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type ResponseList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Data `json:"data"`
}
