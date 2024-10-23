package models

import "picadosYa/internal/entity"

type User struct {
	ID                int64           `json:"id"`
	FirstName         string          `json:"first_name"`
	LastName          string          `json:"last_name"`
	Email             string          `json:"email"`
	Phone             string          `json:"phone"`
	ProfilePictureUrl string          `json:"profile_picture_url"`
	Role              entity.UserRole `json:"role"`
	PositionPlayer    string          `json:"position_player"`
}
