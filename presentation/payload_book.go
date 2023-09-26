package presentation

type BookPayload struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keyword     string   `json:"keyword"`
	Price       float64  `json:"price"`
	Stock       int64    `json:"stock"`
	Publisher   string   `json:"publisher"`
	Categories  []string `json:"categories"`
}

type CategoryPayload struct {
	Category string `json:"category"`
}
