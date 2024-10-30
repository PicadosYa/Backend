package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"picadosYa/encryption"
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

func (s *serv) SavePasswordRecoveryToken(ctx context.Context, email, token string, expiration time.Time) error {
	err := s.repo.SavePasswordRecoveryToken(ctx, email, token, expiration)
	if err != nil {
		return err
	}
	return nil
}

func (s *serv) SendRecoveryEmail(email, token string) error {
	APIKEY := "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"
	baseURL := "http://localhost:8080"

	from := mail.NewEmail("picadosya", "picadosya@gmail.com")
	subject := "Password Recovery"
	to := mail.NewEmail("User", "simonpintos7@gmail.com")
	recoveryURL := fmt.Sprintf("%s/reset-password?token=%s", baseURL, token)
	plainTextContent := fmt.Sprintf("Use the following link to reset your password: %s", recoveryURL)

	htmlContent := fmt.Sprintf("<p>Click <a href='%s'>here</a> to reset your password.</p>", recoveryURL)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

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
