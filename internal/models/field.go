package models

type Field struct {
	Id              int
	Name            string
	Address         string
	Neighborhood    string
	Phone           string
	Latitude        float64
	Longitude       float64
	Type            string
	Price           float64
	Description     string
	LogoUrl         string
	AverageRating   float64
	Services        string
	CreationDate    string
	Photos          []string
	AvailableDays   []Weekday
	UnvailableDates []UnvailableDates
}

type Weekday int

// Esto emula un ENUM, y sirve para especificar que dias de la semana la cancha se abre (EJ: Lunes, Martes, etc)
const (
	Monday Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// Esto sería para fechas no disponibles que el dueño elija (ponele que saco unas vacaciones)
type UnvailableDates struct {
	FromDate string
	ToDate   string
}
