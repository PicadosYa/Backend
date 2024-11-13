package repository

import (
	"context"
	"picadosYa/internal/api/dtos"
	"picadosYa/internal/entity"
	"time"
)

const (
	qryInsertUser = `
	INSERT INTO users (first_name, last_name, email, password, phone, role, accepted_terms)
	VALUES (?, ?, ?, ?, ?, ?, ?);`

	qryGetUserByEmail = `
	select id, first_name, last_name, email, password, phone, profile_picture_url, role, position_player, age, isVerified from users where email = ?;`

	qryInsertToken = `
    INSERT INTO tokens_in_emails (email, token, expires_at)
    VALUES (?, ?, ?);
`

	qryCallGetUserByTokenProcedure = `
	CALL GetUserByToken(?);
	`

	qryUpdateUserStatus = `
	UPDATE users 
	SET isVerified = true 
	WHERE email = ?;
	`

	qryVerifyRecoveryToken = `SELECT COUNT(1) FROM password_recovery_tokens WHERE email = ? AND token = ? AND expires_at > ?`

	qryUpdateUserPassword = `UPDATE users SET password = ? WHERE email = ?`

	qryDeleteRecoveryToken = `DELETE FROM password_recovery_tokens WHERE email = ?`
)

func (r *repo) SaveUser(ctx context.Context, first_name, last_name, email, password, phone string, role entity.UserRole, accepted_terms bool) error {
	// El cifrado de la contraseña va en service
	_, err := r.db.ExecContext(ctx, qryInsertUser, first_name, last_name, email, password, phone, role, accepted_terms)
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

func (r *repo) SaveToken(ctx context.Context, email, token string, expiration time.Time) error {
	_, err := r.db.ExecContext(ctx, qryInsertToken, email, token, expiration)
	return err
}

func (r *repo) GetUserByToken(ctx context.Context, token string) (*dtos.VerifyUserEmail, error) {
	u := &dtos.VerifyUserEmail{}
	err := r.db.GetContext(ctx, u, qryCallGetUserByTokenProcedure, token)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *repo) UpdateUserVerification(ctx context.Context, email string) error {
	_, err := r.db.ExecContext(ctx, qryUpdateUserStatus, email)
	return err
}

func (r *repo) VerifyRecoveryToken(ctx context.Context, email, token string) (bool, error) {
	var count int
	err := r.db.GetContext(ctx, &count, qryVerifyRecoveryToken, email, token, time.Now())
	if err != nil || count == 0 {
		return false, err
	}
	return true, nil
}

func (r *repo) UpdateUserPassword(ctx context.Context, email string, hashedPassword string) error {
	_, err := r.db.ExecContext(ctx, qryUpdateUserPassword, hashedPassword, email)
	return err
}

func (r *repo) DeleteRecoveryToken(ctx context.Context, email string) error {
	_, err := r.db.ExecContext(ctx, qryDeleteRecoveryToken, email)
	return err
}
