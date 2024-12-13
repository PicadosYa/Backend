package repository

import (
	"context"
	"fmt"
	"log"

	"picadosYa/internal/models"

	"github.com/jmoiron/sqlx"
)

type IReservationRepository interface {
	SaveReservation(ctx context.Context, reservation *models.Reservation) error
	GetReservation(ctx context.Context, id int) (*models.Reservation, error)
	GetReservations(ctx context.Context, limit, offset int) ([]models.Reservation, error)
	UpdateReservation(ctx context.Context, reservation *models.Reservation) error
	DeleteReservation(ctx context.Context, id int) error
	GetReservationsPerUser(ctx context.Context, id int) ([]models.Reservations_Result, error)
	//	GetAllReservationsPerOwner(ctx context.Context, id int) ([]models.Reservations_Field_Owner, error)
	GetAllReservationsPerMonth(ctx context.Context, id, month int) ([]models.Reservations_Field_Owner, error)
	// GetAllReservationsPerHour(ctx context.Context, id, hour int) ([]models.Reservations_Field_Owner, error)
	// GetAllReservationsExport(ctx context.Context, id, month, hour int) ([]models.Reservations_Field_Owner, error)
}

type reservationRepository struct {
	db *sqlx.DB
}

func NewReservationRepository(db *sqlx.DB) IReservationRepository {
	return &reservationRepository{
		db: db,
	}
}

func (r *reservationRepository) SaveReservation(ctx context.Context, reservation *models.Reservation) error {
	query := `CALL InsertReservation(?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(
		ctx,
		query,
		reservation.FieldID,
		reservation.UserID,
		reservation.Date,
		reservation.StartTime,
		reservation.EndTime,
		reservation.PaymentID,
	)
	if err != nil {
		log.Fatal("Error executing InsertReservation: ", err)
		return fmt.Errorf("error executing InsertReservation: %w", err)
	}
	return nil
}

func (r *reservationRepository) GetReservationsPerUser(ctx context.Context, id int) ([]models.Reservations_Result, error) {

	query := `CALL GetReservationsByUserId(?)`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []models.Reservations_Result

	for rows.Next() {
		var reservation models.Reservations_Result
		err := rows.Scan(&reservation.EmailUser, &reservation.ReservationDate, &reservation.StartTime, &reservation.EndTime, &reservation.FieldName, &reservation.StatusReservation, &reservation.PaymentID)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

// func (r *reservationRepository) GetAllReservationsExport(ctx context.Context, id, month, hour int) ([]models.Reservations_Field_Owner, error) {
// 	qryGetAllReservations := `CALL GetReservationsPerOwnerByMonthAndHour(?,?,?);`
// 	rows, err := r.db.QueryContext(ctx, qryGetAllReservations, id, month, hour)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var reservations []models.Reservations_Field_Owner

// 	for rows.Next() {
// 		var reservation models.Reservations_Field_Owner
// 		var date string // Usaremos una variable temporal para almacenar la fecha de la base de datos.

// 		err := rows.Scan(&reservation.User_Name, &reservation.Field_Name, &date,
// 			&reservation.Start_Time, &reservation.End_Time, &reservation.Type, &reservation.Phone, &reservation.Status)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Formateamos la fecha a "YYYY-MM-DD".
// 		reservation.Date = date[:10]
// 		reservations = append(reservations, reservation)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return reservations, nil
// }

// func (r *reservationRepository) GetAllReservationsPerOwner(ctx context.Context, id int) ([]models.Reservations_Field_Owner, error) {
// 	qryGetAllReservationsPerFieldOwner := `CALL GetReservationsByOwner(?);`
// 	rows, err := r.db.QueryContext(ctx, qryGetAllReservationsPerFieldOwner, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var reservations []models.Reservations_Field_Owner

// 	for rows.Next() {
// 		var reservation models.Reservations_Field_Owner
// 		var date string // Usaremos una variable temporal para almacenar la fecha de la base de datos.

// 		err := rows.Scan(&reservation.User_Name, &reservation.Field_Name, &date,
// 			&reservation.Start_Time, &reservation.End_Time, &reservation.Type, &reservation.Phone, &reservation.Status)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Formateamos la fecha a "YYYY-MM-DD".
// 		reservation.Date = date[:10]
// 		reservations = append(reservations, reservation)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return reservations, nil
// }

func (r *reservationRepository) GetAllReservationsPerMonth(ctx context.Context, id, month int) ([]models.Reservations_Field_Owner, error) {
	qryGetAllReservationsPerFieldOwner := `CALL GetReservationsPerOwnerByMonth(?, ?);`
	rows, err := r.db.QueryContext(ctx, qryGetAllReservationsPerFieldOwner, id, month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []models.Reservations_Field_Owner

	for rows.Next() {
		var reservation models.Reservations_Field_Owner
		var date string // Usaremos una variable temporal para almacenar la fecha de la base de datos.

		err := rows.Scan(&reservation.User.ID, &reservation.User.Email, &reservation.User.FirstName, &reservation.User.LastName, &reservation.User.ProfilePictureUrl, &reservation.Field.Id, &reservation.Field.Name, &reservation.Field.Address, &reservation.Field.Price, &reservation.Field.Phone, &reservation.Field.Type, &date,
			&reservation.Start_Time, &reservation.End_Time, &reservation.Status)
		if err != nil {
			return nil, err
		}

		// Formateamos la fecha a "YYYY-MM-DD".
		reservation.Date = date[:10] // Aseguramos que solo tomamos los primeros 10 caracteres.

		fmt.Println("reservations per month", reservation)

		reservations = append(reservations, reservation)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

// func (r *reservationRepository) GetAllReservationsPerHour(ctx context.Context, id, hour int) ([]models.Reservations_Field_Owner, error) {
// 	qryGetAllReservationsPerFieldOwner := `CALL GetReservationsPerOwnerByHour(?, ?);`
// 	rows, err := r.db.QueryContext(ctx, qryGetAllReservationsPerFieldOwner, id, hour)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var reservations []models.Reservations_Field_Owner

// 	for rows.Next() {
// 		var reservation models.Reservations_Field_Owner
// 		var date string // Usaremos una variable temporal para almacenar la fecha de la base de datos.

// 		err := rows.Scan(&reservation.User_Name, &reservation.Field_Name, &date,
// 			&reservation.Start_Time, &reservation.End_Time, &reservation.Type, &reservation.Phone, &reservation.Status)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Formateamos la fecha a "YYYY-MM-DD".
// 		reservation.Date = date[:10] // Aseguramos que solo tomamos los primeros 10 caracteres.

// 		fmt.Println("reservations per hour", reservation)

// 		reservations = append(reservations, reservation)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return reservations, nil
// }

func (r *reservationRepository) GetReservation(ctx context.Context, id int) (*models.Reservation, error) {
	query := `CALL GetReservationById(?)`
	var reservation models.Reservation
	err := r.db.GetContext(ctx, &reservation, query, id)
	if err != nil {
		return nil, fmt.Errorf("error executing GetReservationById: %w", err)
	}
	return &reservation, nil
}

func (r *reservationRepository) GetReservations(ctx context.Context, limit, offset int) ([]models.Reservation, error) {
	query := `CALL GetReservationsWithLimitOffset(?, ?)`
	var reservations []models.Reservation
	err := r.db.SelectContext(ctx, &reservations, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error executing GetReservationsWithLimitOffset: %w", err)
	}
	return reservations, nil
}

func (r *reservationRepository) UpdateReservation(ctx context.Context, reservation *models.Reservation) error {
	query := `CALL UpdateReservation(?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(
		ctx,
		query,
		reservation.FieldID,
		reservation.UserID,
		reservation.StartTime,
		reservation.EndTime,
	)
	if err != nil {
		return fmt.Errorf("error executing UpdateReservation: %w", err)
	}
	return nil
}

func (r *reservationRepository) DeleteReservation(ctx context.Context, id int) error {
	query := `CALL DeleteReservation(?)`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error executing DeleteReservation: %w", err)
	}
	return nil
}
