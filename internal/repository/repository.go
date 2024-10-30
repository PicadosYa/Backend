package repository

import (
	"context"

	"picadosYa/internal/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, first_name, last_name, email, password, phone, profile_picture_url string, role entity.UserRole, position_player string, age int) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	SavePasswordRecoveryToken(ctx context.Context, email, token string, expiration time.Time) error
	VerifyRecoveryToken(ctx context.Context, email, token string) (bool, error)
	UpdateUserPassword(ctx context.Context, email string, hashedPassword string) error
	DeleteRecoveryToken(ctx context.Context, email string) error
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
