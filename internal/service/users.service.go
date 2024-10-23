package service

import (
	"context"
	"errors"

	"picadosYa/encryption"
	"picadosYa/internal/entity"
	"picadosYa/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (s *serv) RegisterUser(ctx context.Context, first_name, last_name, email, password, phone, profile_picture_url string, role entity.UserRole, position_player string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	//hash contraseña
	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	pass := encryption.ToBase64(bb)

	return s.repo.SaveUser(ctx, first_name, last_name, email, pass, phone, profile_picture_url, role, position_player)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}
	decryptedPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}
	return &models.User{
		ID:                u.ID,
		FirstName:         u.FirstName,
		LastName:          u.LastName,
		Email:             u.Email,
		Phone:             u.Phone,
		ProfilePictureUrl: u.ProfilePictureUrl,
		Role:              entity.UserRole(u.Role),
		PositionPlayer:    u.PositionPlayer,
	}, nil
}
