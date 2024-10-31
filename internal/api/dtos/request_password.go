package dtos

type RequestPasswordRecovery struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPassword struct {
	Email       string `json:"email" validate:"required,email"`
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}