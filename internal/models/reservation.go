package models

import "time"

type Reservation_without_id struct {
	FieldID   int    `json:"field_id" db:"field_id"`
	Date      string `json:"date" db:"date"`
	StartTime string `json:"start_time" db:"start_time"`
	EndTime   string `json:"end_time" db:"end_time"`
	PaymentID int    `json:"payment_id" db:"payment_id"`
}

type Reservation struct {
	FieldID   int       `json:"field_id" db:"field_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Date      time.Time `json:"date" db:"date"`
	StartTime string    `json:"start_time" db:"start_time"`
	EndTime   string    `json:"end_time" db:"end_time"`
	PaymentID int       `json:"payment_id" db:"payment_id"`
}

type Reservations_Result struct {
	EmailUser         string
	ReservationDate   string
	StartTime         string
	EndTime           string
	FieldName         string
	StatusReservation string
	PaymentID         int
}

type Reservations_Field_Owner struct {
	User_Name  string `json:"user_name"`
	Field_Name string `json:"field_name"`
	Date       string `json:"date"`
	Start_Time string `json:"start_time"`
	End_Time   string `json:"end_time"`
	Type       string `json:"type"`
	Phone      string `json:"phone"`
	Status     string `json:"status"`
}
