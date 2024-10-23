package service

import (
	"context"

	"picadosYa/internal/entity"
	"picadosYa/internal/models"
	"picadosYa/internal/repository"
)

// Logica de la aplicacion
//
//go:generate mockery --name=Service --output:service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, first_name, last_name, email, password, phone, profile_picture_url string, role entity.UserRole, position_player string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
