package presentation

type ResponseRegister struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    DataRegister `json:"data"`
}

type DataRegister struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
}
