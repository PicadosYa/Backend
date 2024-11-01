package service

import (
	"context"
	"picadosYa/internal/entity"
	"picadosYa/internal/models"
	"picadosYa/internal/repository"
	"time"
)

// Logica de la aplicacion
//
//go:generate mockery --name=Service --output:service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, first_name, last_name, email, password, phone, profile_picture_url string, role entity.UserRole, position_player string, age int) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	SavePasswordRecoveryToken(ctx context.Context, email, token string, expiration time.Time) error
	SendRecoveryEmail(email, token string) error
	ResetPassword(ctx context.Context, email, token, newPassword string) error
	VerifyRecoveryToken(ctx context.Context, email, token string) (bool, error)
	DeleteRecoveryToken(ctx context.Context, email string) error
	UpdateUserPassword(ctx context.Context, email string, hashedPassword string) error
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
