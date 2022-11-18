package response

type ErrorRes struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type SuccessRes struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}