package dtos

import "picadosYa/internal/entity"

type RegisterUser struct {
	FirstName     string          `json:"first_name" validate:"required"`
	Lastname      string          `json:"last_name" validate:"required"`
	Email         string          `json:"email" validate:"required,email"`
	Password      string          `json:"password" validate:"required,min=8"`
	Phone         string          `json:"phone" validate:"required"`
	Role          entity.UserRole `json:"role" validate:"required"`
	AcceptedTerms bool            `json:"accepted_terms" validate:"required"`
}

type RegisteredUser struct {
	FirstName string          `json:"first_name" validate:"required"`
	LastName  string          `json:"last_name" validate:"required"`
	Email     string          `json:"email" validate:"required,email"`
	Phone     string          `json:"phone" validate:"required"`
	Role      entity.UserRole `json:"role" validate:"required"`
}

type LoguedUser struct {
	FirstName  string          `json:"first_name" validate:"required"`
	LastName   string          `json:"last_name" validate:"required"`
	Email      string          `json:"email" validate:"required,email"`
	Phone      string          `json:"phone" validate:"required"`
	Role       entity.UserRole `json:"role" validate:"required"`
	IsVerified bool            `json:"isVerified"`
}

type UpdateUser struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	PositionPlayer    string `json:"position_player"`
	TeamName          string `json:"team_name"`
	Age               int    `json:"age"`
	ProfilePictureUrl string `json:"profile_picture_url"`
	ID                int    `json:"id"`
}
