package presentation

type ResponseUpdate struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    DataUpdate `json:"data"`
}

type DataUpdate struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    []string `json:"category"`
	Keyword     string   `json:"keyword"`
	Price       string   `json:"price"`
	Stock       int64    `json:"stock"`
	Publisher   string   `json:"publisher"`
	UpdatedAt   string   `json:"updated_at"`
}
