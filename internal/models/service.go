package models

type Service struct {
	ID   int    `json:"id" db:"id" form:"id"`
	Name string `json:"name" db:"name"`
	Icon string `json:"icon" db:"icon"`
}
