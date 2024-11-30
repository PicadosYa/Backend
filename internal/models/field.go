package models

import (
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
)

type Field struct {
	Id              int                  `json:"id" db:"id" form:"id"`
	Name            string               `json:"name" db:"name" form:"name"`
	Address         string               `json:"address" db:"address" form:"address"`
	Neighborhood    string               `json:"neighborhood" db:"neighborhood" form:"neighborhood"`
	Phone           string               `json:"phone" db:"phone" form:"phone"`
	Latitude        float64              `json:"latitude" db:"latitude" form:"latitude"`
	Longitude       float64              `json:"longitude" db:"longitude" form:"longitude"`
	Type            string               `json:"type" db:"type" default:"5" form:"type"`
	Price           float64              `json:"price" db:"price" form:"price"`
	Description     string               `json:"description" db:"description" form:"description"`
	LogoUrl         string               `json:"logo_url" db:"logo_url" form:"logo_url"`
	AverageRating   float64              `json:"average_rating" db:"average_rating" form:"average_rating"`
	Services        []Service            `json:"services" db:"services" form:"services"`
	CreationDate    strfmt.Date          `json:"creation_date" db:"creation_date" form:"creation_date"`
	Photos          []string             `json:"photos" db:"photos" form:"photos"`
	AvailableDays   []string             `json:"available_days" db:"available_days" form:"available_days"`
	UnvailableDates []UnvailableDates    `json:"unvailable_dates" db:"unvailable_dates" form:"unvailable_dates"`
	Reservations    []ReservationReduced `json:"reservations" db:"reservations" form:"reservations"`
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

type FieldWithID_User struct {
	Id              int                  `json:"id" db:"id" form:"id"`
	Name            string               `json:"name" db:"name" form:"name"`
	Address         string               `json:"address" db:"address" form:"address"`
	Neighborhood    string               `json:"neighborhood" db:"neighborhood" form:"neighborhood"`
	Phone           string               `json:"phone" db:"phone" form:"phone"`
	Latitude        float64              `json:"latitude" db:"latitude" form:"latitude"`
	Longitude       float64              `json:"longitude" db:"longitude" form:"longitude"`
	Type            string               `json:"type" db:"type" default:"5" form:"type"`
	Price           float64              `json:"price" db:"price" form:"price"`
	Description     string               `json:"description" db:"description" form:"description"`
	LogoUrl         string               `json:"logo_url" db:"logo_url" form:"logo_url"`
	AverageRating   float64              `json:"average_rating" db:"average_rating" form:"average_rating"`
	Services        []Service            `json:"services" db:"services" form:"services"`
	CreationDate    strfmt.Date          `json:"creation_date" db:"creation_date" form:"creation_date"`
	Photos          []string             `json:"photos" db:"photos" form:"photos"`
	AvailableDays   []string             `json:"available_days" db:"available_days" form:"available_days"`
	UnvailableDates []UnvailableDates    `json:"unvailable_dates" db:"unvailable_dates" form:"unvailable_dates"`
	Reservations    []ReservationReduced `json:"reservations" db:"reservations" form:"reservations"`
	ID_User         int                  `json:"user_id" db:"user_id"`
}

type FieldsResultsPerOwner struct {
	Field_Name    string `json:"field_name"`
	Field_Address string `json:"field_address"`
	Field_Type    string `json:"field_type"`
	Field_Phone   string `json:"field_phone"`
	Field_Status  bool   `json:"field_status"`
}
