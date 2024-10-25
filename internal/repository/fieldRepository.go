package repository

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"picadosYa/internal/models"
	"picadosYa/utils"

	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
)

type IFieldRepository interface {
	SaveField(ctx context.Context, field *models.Field) error
	GetField(ctx context.Context, id int, month time.Time) (*models.Field, error)
	GetFields(ctx context.Context) ([]models.Field, error)
	UpdateField(ctx context.Context, field *models.Field) error
	RemoveField(ctx context.Context, id int) error
}

type fieldRepository struct {
	db *sqlx.DB
}

func NewFieldRepository(db *sqlx.DB) IFieldRepository {
	return &fieldRepository{
		db: db,
	}
}

func (r *fieldRepository) SaveField(ctx context.Context, field *models.Field) error {
	query := `CALL InsertField(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	photoURLsStr := strings.Join(field.Photos, ",")
	availableDaysStr := strings.Join(field.AvailableDays, ",")
	serviceIDsStr := utils.SliceToString(field.Services)

	if field.Type == "" {
		field.Type = "5"
	}

	if field.CreationDate.String() == "0001-01-01" {
		field.CreationDate = strfmt.Date(time.Now())
	}

	log.Println("Query: ", query)
	log.Println("Type: ", field.Type)
	// Ejecutar el procedimiento almacenado con los parámetros del campo
	_, err := r.db.ExecContext(
		ctx,
		query,
		field.Name,
		field.Address,
		field.Neighborhood,
		field.Phone,
		field.Latitude,
		field.Longitude,
		field.Type,
		field.Price,
		field.Description,
		field.LogoUrl,
		field.CreationDate,
		availableDaysStr,
		photoURLsStr,
		serviceIDsStr,
	)

	if err != nil {
		log.Fatal("Error executing InsertField: ", err)
		return fmt.Errorf("error executing InsertField: %w", err)
	}

	return nil
}

func (r *fieldRepository) GetField(ctx context.Context, id int, month time.Time) (*models.Field, error) {
	// Definir la consulta SQL con el procedimiento almacenado
	query := `CALL GetFieldReservationsByMonthAndId(?, ?)`

	// Crear una variable para almacenar el resultado del campo
	var field models.Field

	// Ejecutar la consulta
	rows, err := r.db.QueryContext(ctx, query, id, month.Format("2006-01"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			photos           string
			services         string
			unavailableDates string
			reservations     string
			availableDays    string
		)

		// Escanear las columnas del resultado en las variables correspondientes
		err := rows.Scan(
			&field.Id,
			&field.Name,
			&field.Address,
			&field.Neighborhood,
			&field.Phone,
			&field.Latitude,
			&field.Longitude,
			&field.Type,
			&field.Price,
			&field.Description,
			&field.LogoUrl,
			&field.AverageRating,
			&field.CreationDate,
			&availableDays,    // Cadena con los días disponibles
			&photos,           // Cadena con las URLs de fotos
			&services,         // Cadena con los nombres de servicios
			&unavailableDates, // Cadena con las fechas no disponibles
			&reservations,     // Cadena con las reservas
		)
		if err != nil {
			return nil, err
		}

		// Parsear los días disponibles, fotos, servicios, fechas no disponibles y reservas
		field.AvailableDays = utils.SplitString(availableDays)
		field.Photos = utils.SplitString(photos)
		field.Services = utils.ParseServices(services)
		field.UnvailableDates = utils.ParseUnavailableDates(unavailableDates)
		field.Reservations = utils.ParseReservations(reservations)
	}

	// Verificar si hubo algún error durante el proceso
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &field, nil
}

func (r *fieldRepository) GetFields(ctx context.Context) ([]models.Field, error) {
	// TODO: Implementar
	return nil, nil
}

func (r *fieldRepository) UpdateField(ctx context.Context, field *models.Field) error {
	// TODO: Implementar
	return nil
}

func (r *fieldRepository) RemoveField(ctx context.Context, id int) error {
	// TODO: Implementar
	return nil
}
