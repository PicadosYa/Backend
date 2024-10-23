package entity

type User struct {
	ID            int64  `db:"id"`
	Email         string `db:"email"`
	Name          string `db:"name"`
	Password      string `db:"password"`
	Lastname      string `db:"lastname"`
	Telephone     string `db:"telephone"`
	Profile_photo string `db:"profile_photo"`
}
