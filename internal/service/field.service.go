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
	GetFieldsPerOwner(ctx context.Context, id_user int) ([]models.FieldsResultsPerOwner, error)
	GetFieldIndividually(ctx context.Context, id int) *models.FieldsReduced
}

type fieldService struct {
	repo     repository.IFieldRepository
	fileRepo repository.IFileRepository
}

func NewFieldService(repo repository.IFieldRepository, fileRepo repository.IFileRepository) FieldService {
	return &fieldService{
		repo:     repo,
		fileRepo: fileRepo,
	}
}

func (s *fieldService) SaveField(ctx context.Context, field *models.Field, files *map[string][]*multipart.FileHeader) error {
	log.Println("Saving field")
	log.Printf("Files: %v", files)
	for key, fileHeaders := range *files {
		if key == "fieldImages" {
			log.Println(key)
			log.Printf("fileHeaders: %v", fileHeaders)
			for _, fileHeader := range fileHeaders {
				log.Printf("fileHeader: %s", fileHeader.Filename)
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

func (s *fieldService) GetFieldsPerOwner(ctx context.Context, id_user int) ([]models.FieldsResultsPerOwner, error) {
	return s.repo.GetFieldsPerOwner(ctx, id_user)
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

func (s *fieldService) GetFieldIndividually(ctx context.Context, id int) *models.FieldsReduced {
	return s.repo.GetFieldIndividually(ctx, id)
}
