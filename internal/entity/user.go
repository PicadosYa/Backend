package entity

type User struct {
	ID                int64    `db:"id"`
	FirstName         string   `db:"first_name"`
	LastName          string   `db:"last_name"`
	Email             string   `db:"email"`
	Password          string   `db:"password"`
	Phone             string   `db:"phone"`
	ProfilePictureUrl string   `db:"profile_picture_url"`
	Role              UserRole `db:"role"`
	PositionPlayer    string   `db:"position_player"`
	Age               int      `db:"age"`
	IsVerified        bool     `db:"isVerified"`
	AcceptedTerms     bool     `db:"accepted_terms"`
}

type UserEmailByID struct {
	Email string `json:"user_email"`
}
