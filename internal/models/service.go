package models

type Service struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Icon string `json:"icon" db:"icon"`
}
