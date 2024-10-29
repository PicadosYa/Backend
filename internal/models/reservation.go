package models

import "time"

type Reservation struct {
	Id              int       `json:"id" db:"id"`
	FieldID         int       `json:"field_id" db:"field_id"`
	UserID          int       `json:"user_id" db:"user_id"`
	Date            time.Time `json:"date" db:"date"`
	StartTime       []uint8   `json:"start_time" db:"start_time"`
	EndTime         []uint8   `json:"end_time" db:"end_time"`
	Status          string    `json:"status" db:"status"`
	ReservationDate time.Time `json:"reservation_date" db:"reservation_date"`
}
