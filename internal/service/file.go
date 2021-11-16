package service

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/Alexander272/astro-atlas/internal/models"
	"github.com/Alexander272/astro-atlas/pkg/logger"
	"github.com/Alexander272/astro-atlas/pkg/storage"
)

type FileService struct {
	storage storage.Provider
}

func NewFileService(storage storage.Provider) *FileService {
	return &FileService{
		storage: storage,
	}
}

func (s *FileService) Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader, path, filename string) (models.File, error) {
	res, err := s.storage.Upload(ctx, file, header, path, filename)
	if err != nil {
		return models.File{}, fmt.Errorf("failed to upload file. error: %w", err)
	}

	logger.Debug(header.Header)
	uploadedFile := models.File{
		FileType: header.Header.Get("Content-Type"),
		Name:     res.Name,
		OrigName: header.Filename,
		Url:      res.Url,
	}
	return uploadedFile, nil
}

func (s *FileService) Delete(ctx context.Context, path, filename string) error {
	logger.Debug(filename)
	if err := s.storage.Remove(ctx, path, filename); err != nil {
		return fmt.Errorf("failed to delete file. error: %w", err)
	}
	return nil
}
