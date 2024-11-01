package repository

import (
	"context"
	"picadosYa/internal/entity"
	"time"
)

const (
	qryInsertUser = `
	INSERT INTO users (first_name, last_name, email, password, phone, profile_picture_url, role, position_player, age)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`

	qryGetUserByEmail = `
	select id, first_name, last_name, email, password, phone, profile_picture_url, role, position_player, age from users where email = ?;`

	qryInsertRecoveryToken = `
    INSERT INTO password_recovery_tokens (email, token, expires_at)
    VALUES (?, ?, ?);
`
)

func (r *repo) SaveUser(ctx context.Context, first_name, last_name, email, password, phone, profile_picture_url string, role entity.UserRole, position_player string, age int) error {
	// El cifrado de la contraseña va en service
	_, err := r.db.ExecContext(ctx, qryInsertUser, first_name, last_name, email, password, phone, profile_picture_url, role, position_player, age)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, qryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *repo) SavePasswordRecoveryToken(ctx context.Context, email, token string, expiration time.Time) error {
	_, err := r.db.ExecContext(ctx, qryInsertRecoveryToken, email, token, expiration)
	return err
}

func (r *repo) VerifyRecoveryToken(ctx context.Context, email, token string) (bool, error) {
	var count int
	query := `SELECT COUNT(1) FROM password_recovery_tokens WHERE email = ? AND token = ? AND expires_at > ?`
	err := r.db.GetContext(ctx, &count, query, email, token, time.Now())
	if err != nil || count == 0 {
		return false, err
	}
	return true, nil
}
func (r *repo) UpdateUserPassword(ctx context.Context, email string, hashedPassword string) error {
	query := `UPDATE users SET password = ? WHERE email = ?`
	_, err := r.db.ExecContext(ctx, query, hashedPassword, email)
	return err
}

func (r *repo) DeleteRecoveryToken(ctx context.Context, email string) error {
	query := `DELETE FROM password_recovery_tokens WHERE email = ?`
	_, err := r.db.ExecContext(ctx, query, email)
	return err
}
