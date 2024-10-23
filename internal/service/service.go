package service

import (
	"context"

	"picadosYa/internal/models"
	"picadosYa/internal/repository"
)

// Logica de la aplicacion
//
//go:generate mockery --name=Service --output:service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, lastname, password, telephone, profile_photo string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	AddUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
