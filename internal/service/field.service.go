package service

import (
	"context"
	"log"
	"picadosYa/internal/models"
	"picadosYa/internal/repository"
	"time"
)

type FieldService interface {
	SaveField(ctx context.Context, field *models.Field) error
	GetField(ctx context.Context, id int, month time.Time) (*models.Field, error)
	GetFields(ctx context.Context) ([]models.Field, error)
	UpdateField(ctx context.Context, field *models.Field) error
	RemoveField(ctx context.Context, id int) error
}

type fieldService struct {
	repo repository.IFieldRepository
}

func NewFieldService(repo repository.IFieldRepository) FieldService {
	return &fieldService{
		repo: repo,
	}
}

func (s *fieldService) SaveField(ctx context.Context, field *models.Field) error {
	log.Println("Saving field")
	return s.repo.SaveField(ctx, field)
}

func (s *fieldService) GetField(ctx context.Context, id int, month time.Time) (*models.Field, error) {
	return s.repo.GetField(ctx, id, month)
}

func (s *fieldService) GetFields(ctx context.Context) ([]models.Field, error) {
	return s.repo.GetFields(ctx)
}

func (s *fieldService) UpdateField(ctx context.Context, field *models.Field) error {
	return s.repo.UpdateField(ctx, field)
}

func (s *fieldService) RemoveField(ctx context.Context, id int) error {
	return s.repo.RemoveField(ctx, id)
}
