package models

type Service struct {
	ID   int    `json:"id" db:"id" form:"id"`
	Name string `json:"name" db:"name"`
	Icon string `json:"icon" db:"icon"`
}

type ResponseMessage struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ResponseError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
