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
