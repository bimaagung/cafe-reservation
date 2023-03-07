package response

type ErrorRes struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type SuccessRes struct {
	Status  string      `json:"status" example:"ok"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}