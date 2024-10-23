package dtos

type RegisterUser struct {
	Email         string `json:"email" validate:"required,email"`
	Name          string `json:"name" validate:"required"`
	Lastname      string `json:"lastname" validate:"required"`
	Password      string `json:"password" validate:"required,min=8"`
	Telephone     string `json:"telephone" validate:"required"`
	Profile_photo string `json:"profile_photo" validate:"required"`
}
