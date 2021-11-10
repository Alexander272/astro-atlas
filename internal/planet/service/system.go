package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alexander272/astro-atlas/internal/planet/models"
	"github.com/Alexander272/astro-atlas/internal/repository"
	"github.com/Alexander272/astro-atlas/pkg/apperror"
)

type SystemService struct {
	repo *repository.Repositories
}

func NewSystemService(repo *repository.Repositories) *SystemService {
	return &SystemService{
		repo: repo,
	}
}

func (s *SystemService) Create(ctx context.Context, dto models.CreateSystemDTO) (systemId string, err error) {
	system := models.NewSystem(dto)
	systemId, err = s.repo.System.Create(ctx, system)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return systemId, err
		}
		return systemId, fmt.Errorf("failed to create system. error: %w", err)
	}

	return systemId, nil
}

func (s *SystemService) GetList(ctx context.Context) ([]models.SystemShort, error) {
	systems, err := s.repo.System.GetList(ctx)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return systems, err
		}
		return systems, fmt.Errorf("failed to get systems. error: %w", err)
	}
	if len(systems) == 0 {
		return systems, apperror.ErrNotFound
	}

	return systems, nil
}

func (s *SystemService) GetById(ctx context.Context, systemId string) (models.System, error) {
	system, err := s.repo.System.GetById(ctx, systemId)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return system, err
		}
		return system, fmt.Errorf("failed to find system by id. error: %w", err)
	}

	return system, nil
}

func (s *SystemService) Update(ctx context.Context, dto models.UpdateSystemDTO) error {
	err := s.repo.System.Update(ctx, models.System(dto))
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return err
		}
		return fmt.Errorf("failed to update system. error: %w", err)
	}
	return nil
}

func (s *SystemService) Delete(ctx context.Context, systemId string) error {
	err := s.repo.System.Delete(ctx, systemId)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return err
		}
		return fmt.Errorf("failed to delete system. error: %w", err)
	}
	return nil
}
