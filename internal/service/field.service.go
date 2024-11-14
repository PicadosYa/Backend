package service

import (
	"context"
	"log"
	"mime/multipart"
	"picadosYa/internal/models"
	"picadosYa/internal/repository"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type FieldService interface {
	SaveField(ctx context.Context, field *models.Field, files *map[string][]*multipart.FileHeader) error
	GetField(ctx context.Context, id int, month time.Time) (*models.Field, error)
	GetFields(ctx context.Context, month time.Time, limit int, offset int) ([]models.Field, error)
	UpdateField(ctx context.Context, field *models.Field) error
	PatchField(ctx context.Context, field *models.Field) error
	RemoveField(ctx context.Context, id int) error
}

type fieldService struct {
	repo     repository.IFieldRepository
	fileRepo repository.IFileRepository
}

func NewFieldService(repo repository.IFieldRepository) FieldService {
	return &fieldService{
		repo: repo,
	}
}

func (s *fieldService) SaveField(ctx context.Context, field *models.Field, files *map[string][]*multipart.FileHeader) error {
	log.Println("Saving field")
	for key, fileHeaders := range *files {
		if key == "fieldImages" {
			for _, fileHeader := range fileHeaders {
				photo, err := s.fileRepo.UploadFile(fileHeader, fileHeader.Filename+"_"+uuid.New().String())
				if err != nil {
					return err
				}
				field.Photos = append(field.Photos, photo)
			}
		}
	}
	return s.repo.SaveField(ctx, field)
}

func (s *fieldService) GetField(ctx context.Context, id int, month time.Time) (*models.Field, error) {
	return s.repo.GetField(ctx, id, month)
}

func (s *fieldService) GetFields(ctx context.Context, month time.Time, limit int, offset int) ([]models.Field, error) {
	return s.repo.GetFields(ctx, month, limit, offset)
}

func (s *fieldService) UpdateField(ctx context.Context, field *models.Field) error {
	err := validation.ValidateStruct(field,
		validation.Field(&field.Id, validation.Required),
	)

	if err != nil {
		return err
	}
	return s.repo.UpdateField(ctx, field)
}

func (s *fieldService) PatchField(ctx context.Context, field *models.Field) error {
	err := validation.ValidateStruct(field,
		validation.Field(&field.Id, validation.Required),
	)

	if err != nil {
		return err
	}
	return s.repo.PatchField(ctx, field)
}

func (s *fieldService) RemoveField(ctx context.Context, id int) error {
	return s.repo.RemoveField(ctx, id)
}
