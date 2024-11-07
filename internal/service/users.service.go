package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"picadosYa/encryption"
	"picadosYa/internal/api/dtos"
	"picadosYa/internal/entity"
	"picadosYa/internal/models"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrTokenInvalidOrExpired = errors.New("invalid or expired token")
)

const APIKEY = "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"

func (s *serv) RegisterUser(ctx context.Context, first_name, last_name, email, password, phone, profile_picture_url string, role entity.UserRole, position_player string, edad int) error {
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

	return s.repo.SaveUser(ctx, first_name, last_name, email, pass, phone, profile_picture_url, role, position_player, edad)
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
		Age:               u.Age,
	}, nil
}

func (s *serv) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *serv) SaveToken(ctx context.Context, email, token string, expiration time.Time) error {
	err := s.repo.SaveToken(ctx, email, token, expiration)
	if err != nil {
		return err
	}
	return nil
}

func (s *serv) SendRecoveryEmail(email, token string) error {
	ctx := context.Background()
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil
	}
	templateID := "d-14d7497e32d745889c502d5bb3d7bdca"
	return sendEmail(templateID, email, token, u.FirstName)
}

func (s *serv) ResetPassword(ctx context.Context, email, token, newPassword string) error {
	//Verifica el token
	valid, err := s.repo.VerifyRecoveryToken(ctx, email, token)
	if err != nil || !valid {
		return ErrTokenInvalidOrExpired
	}

	bb, err := encryption.Encrypt([]byte(newPassword))
	if err != nil {
		return err
	}
	pass := encryption.ToBase64(bb)

	err = s.repo.UpdateUserPassword(ctx, email, pass)
	if err != nil {
		log.Println("No está actualizando la contraseña")
		return err
	}

	// Eliminar el token de recuperación después de su uso
	err = s.repo.DeleteRecoveryToken(ctx, email)
	if err != nil {
		log.Println("No está eliminando el token")
		return err
	}
	return nil
}

func (s *serv) DeleteRecoveryToken(ctx context.Context, email string) error {
	return s.repo.DeleteRecoveryToken(ctx, email)
}
func (s *serv) UpdateUserPassword(ctx context.Context, email string, hashedPassword string) error {
	return s.repo.UpdateUserPassword(ctx, email, hashedPassword)
}
func (s *serv) VerifyRecoveryToken(ctx context.Context, email, token string) (bool, error) {
	return s.repo.VerifyRecoveryToken(ctx, email, token)
}

func (s *serv) SendVerifyEmail(email, token string) error {
	ctx := context.Background()
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil
	}
	templateID := "d-b512ab2466914e5fb4315a7e0998506c"
	return sendEmail(templateID, email, token, u.FirstName)
}

func (s *serv) GetUserByToken(ctx context.Context, token string) (*dtos.VerifyUserEmail, error) {
	return s.repo.GetUserByToken(ctx, token)
}

func (s *serv) UpdateUserVerification(ctx context.Context, email string) error {
	return s.repo.UpdateUserVerification(ctx, email)
}

func sendEmail(templateID, email, token, name string) error {
	message := mail.NewV3Mail()
	from := mail.NewEmail("picadosya", "picadosya@gmail.com")
	message.SetFrom(from)
	personalization := mail.NewPersonalization()
	to := mail.NewEmail("picadosya", email)
	personalization.AddTos(to)
	personalization.SetDynamicTemplateData("name", name)
	personalization.SetDynamicTemplateData("token", token)
	message.AddPersonalizations(personalization)
	message.SetTemplateID(templateID)
	client := sendgrid.NewSendClient(APIKEY)
	response, err := client.Send(message)
	if err != nil {
		return err
	}
	if response.StatusCode != 202 {
		return fmt.Errorf("failed to send email, status code: %d", response.StatusCode)
	}
	return nil
}
