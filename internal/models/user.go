package models

type User struct {
	ID            int64  `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Lastname      string `json:"lastname"`
	Telephone     string `json:"telephone"`
	Profile_photo string `json:"profile_photo"`
}
