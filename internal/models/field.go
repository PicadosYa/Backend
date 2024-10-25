package models

import (
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
)

type Field struct {
	Id              int                  `json:"id" db:"id"`
	Name            string               `json:"name" db:"name"`
	Address         string               `json:"address" db:"address"`
	Neighborhood    string               `json:"neighborhood" db:"neighborhood"`
	Phone           string               `json:"phone" db:"phone"`
	Latitude        float64              `json:"latitude" db:"latitude"`
	Longitude       float64              `json:"longitude" db:"longitude"`
	Type            string               `json:"type" db:"type" default:"5"`
	Price           float64              `json:"price" db:"price"`
	Description     string               `json:"description" db:"description"`
	LogoUrl         string               `json:"logo_url" db:"logo_url"`
	AverageRating   float64              `json:"average_rating" db:"average_rating"`
	Services        []Service            `json:"services" db:"services"`
	CreationDate    strfmt.Date          `json:"creation_date" db:"creation_date" default:"now()"`
	Photos          []string             `json:"photos" db:"photos"`
	AvailableDays   []string             `json:"available_days" db:"available_days"`
	UnvailableDates []UnvailableDates    `json:"unvailable_dates" db:"unvailable_dates"`
	Reservations    []ReservationReduced `json:"reservations" db:"reservations"`
}

// Esto sería para fechas no disponibles que el dueño elija (ponele que saco unas vacaciones)
type UnvailableDates struct {
	FromDate string
	ToDate   string
}

type ReservationReduced struct {
	Date      strfmt.Date
	StartTime HourMinute
	EndTime   HourMinute
}

type HourMinute time.Time

// Personalizar la salida JSON para el tipo HourMinute
func (hm HourMinute) MarshalJSON() ([]byte, error) {
	// Formatear la hora en "15:04" (hora:minutos)
	formatted := fmt.Sprintf("\"%s\"", time.Time(hm).Format("15:04"))
	return []byte(formatted), nil
}
