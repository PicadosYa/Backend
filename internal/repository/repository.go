package repository

import (
	"context"

	"picadosYa/internal/api/dtos"
	"picadosYa/internal/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, first_name, last_name, email, password, phone string, role entity.UserRole, accepted_terms bool) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	SaveToken(ctx context.Context, email, token string, expiration time.Time) error
	GetUserByToken(ctx context.Context, token string) (*dtos.VerifyUserEmail, error)
	UpdateUserVerification(ctx context.Context, email string) error
	VerifyRecoveryToken(ctx context.Context, email, token string) (bool, error)
	UpdateUserPassword(ctx context.Context, email string, hashedPassword string) error
	UpdateUserProfileInfo(ctx context.Context, first_name, last_name, email, phone, position_player, team_name string, age int, profile_picture_url string, id int) error
	DeleteRecoveryToken(ctx context.Context, email string) error
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	CreateOrRemoveFavourite(ctx context.Context, id_user, id_field int) error
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
