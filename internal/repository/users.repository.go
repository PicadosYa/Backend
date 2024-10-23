package repository

import (
	"context"

	"picadosYa/internal/entity"
)

const (
	qryInsertUser = `
	INSERT INTO USERS (first_name, last_name, email, password, phone, profile_picture_url, role, position_player)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	qryGetUserByEmail = `
	select id, first_name, last_name, email, password, phone, profile_picture_url, role, position_player from USERS where email = ?;`
)

func (r *repo) SaveUser(ctx context.Context, first_name, last_name, email, password, phone, profile_picture_url string, role entity.UserRole, position_player string) error {
	// El cifrado de la contraseña va en service
	_, err := r.db.ExecContext(ctx, qryInsertUser, first_name, last_name, email, password, phone, profile_picture_url, role, position_player)
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
