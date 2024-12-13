package service

import (
	"context"
	"log"
	"picadosYa/internal/models"
	"picadosYa/internal/repository"
)

type ReservationService interface {
	SaveReservation(ctx context.Context, reservation *models.Reservation) error
	GetReservation(ctx context.Context, id int) (*models.Reservation, error)
	GetReservations(ctx context.Context, limit, offset int) ([]models.Reservation, error)
	CreateReservation(ctx context.Context, reservation *models.Reservation) error
	UpdateReservation(ctx context.Context, reservation *models.Reservation) error
	DeleteReservation(ctx context.Context, id int) error
	GetReservationsPerUser(ctx context.Context, id int) ([]models.Reservations_Result, error)
	//GetAllReservationsPerOwner(ctx context.Context, id int) ([]models.Reservations_Field_Owner, error)
	GetAllReservationsPerMonth(ctx context.Context, id, month int) ([]models.Reservations_Field_Owner, error)
	//GetAllReservationsPerHour(ctx context.Context, id, hour int) ([]models.Reservations_Field_Owner, error)
	GetAllReservationsExport(ctx context.Context, id, month, hour int) ([]models.Reservations_Field_Owner, error)
}

type reservationService struct {
	repo repository.IReservationRepository
}

func NewReservationService(repo repository.IReservationRepository) ReservationService {
	return &reservationService{
		repo: repo,
	}
}

func (s *reservationService) SaveReservation(ctx context.Context, reservation *models.Reservation) error {
	log.Println("Saving reservation")
	return s.repo.SaveReservation(ctx, reservation)
}

func (s *reservationService) GetAllReservationsExport(ctx context.Context, id, month, hour int) ([]models.Reservations_Field_Owner, error) {
	if hour == 123 && month == 123 {
		return s.repo.GetAllReservationsPerMonth(ctx, id, 12)
	} else if month == 123 {
		//	return s.repo.GetAllReservationsPerHour(ctx, id, hour)
	} else if hour == 123 {
		return s.repo.GetAllReservationsPerMonth(ctx, id, month)
	}
	return s.repo.GetAllReservationsPerMonth(ctx, id, month)
}

func (s *reservationService) GetReservationsPerUser(ctx context.Context, id int) ([]models.Reservations_Result, error) {
	return s.repo.GetReservationsPerUser(ctx, id)
}

func (s *reservationService) GetReservation(ctx context.Context, id int) (*models.Reservation, error) {
	return s.repo.GetReservation(ctx, id)
}

func (s *reservationService) GetReservations(ctx context.Context, limit, offset int) ([]models.Reservation, error) {
	return s.repo.GetReservations(ctx, limit, offset)
}

func (s *reservationService) CreateReservation(ctx context.Context, reservation *models.Reservation) error {
	log.Println("Creating reservation")
	return s.repo.SaveReservation(ctx, reservation)
}

func (s *reservationService) UpdateReservation(ctx context.Context, reservation *models.Reservation) error {
	return s.repo.UpdateReservation(ctx, reservation)
}

func (s *reservationService) DeleteReservation(ctx context.Context, id int) error {
	return s.repo.DeleteReservation(ctx, id)
}

//	func (s *reservationService) GetAllReservationsPerOwner(ctx context.Context, id int) ([]models.Reservations_Field_Owner, error) {
//		return s.repo.GetAllReservationsPerOwner(ctx, id)
//	}
func (s *reservationService) GetAllReservationsPerMonth(ctx context.Context, id, month int) ([]models.Reservations_Field_Owner, error) {
	return s.repo.GetAllReservationsPerMonth(ctx, id, month)
}

// func (s *reservationService) GetAllReservationsPerHour(ctx context.Context, id, hour int) ([]models.Reservations_Field_Owner, error) {
// 	return s.repo.GetAllReservationsPerHour(ctx, id, hour)
// }
