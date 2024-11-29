package service

import (
	"context"
	"mime/multipart"
	"picadosYa/internal/api/dtos"
	"picadosYa/internal/entity"
	"picadosYa/internal/models"
	"picadosYa/internal/repository"
	"time"
)

// Logica de la aplicacion
//
//go:generate mockery --name=Service --output:service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, first_name, last_name, email, password, phone string, role entity.UserRole, accepted_terms bool) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	SaveToken(ctx context.Context, email, token string, expiration time.Time) error
	GetUserByToken(ctx context.Context, token string) (*dtos.VerifyUserEmail, error)
	UpdateUserVerification(ctx context.Context, email string) error
	SendRecoveryEmail(email, token string) error
	SendVerifyEmail(email, token string) error
	ResetPassword(ctx context.Context, email, token, newPassword string) error
	VerifyRecoveryToken(ctx context.Context, email, token string) (bool, error)
	DeleteRecoveryToken(ctx context.Context, email string) error
	UpdateUserPassword(ctx context.Context, email string, hashedPassword string) error
	UpdateUserInfo(ctx context.Context, first_name, last_name, email, phone, position_player, team_name string, age int, file *multipart.FileHeader, id int) (string, error)
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	CreateOrRemoveFavourite(ctx context.Context, id_user, id_field int) error
	GetFavouritesPerUser(ctx context.Context, id int) ([]dtos.FavsResults, error)
}

type serv struct {
	repo     repository.Repository
	fileRepo repository.IFileRepository
}

func New(repo repository.Repository, fileRepo repository.IFileRepository) Service {
	return &serv{
		repo:     repo,
		fileRepo: fileRepo,
	}
}
