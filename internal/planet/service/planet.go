package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alexander272/astro-atlas/internal/planet/models"
	"github.com/Alexander272/astro-atlas/internal/repository"
	"github.com/Alexander272/astro-atlas/pkg/apperror"
)

type PlanetService struct {
	repo *repository.Repositories
}

func NewPlanetService(repo *repository.Repositories) *PlanetService {
	return &PlanetService{
		repo: repo,
	}
}

func (s *PlanetService) Create(ctx context.Context, dto models.CreatePlanetDTO) (string, error) {
	return "", nil
}

func (s *PlanetService) GetList(ctx context.Context, systemId string) ([]models.PlanetShort, error) {
	planets, err := s.repo.Planet.GetList(ctx, systemId)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return planets, err
		}
		return planets, fmt.Errorf("failed to get planets by systemid. error: %w", err)
	}
	if len(planets) == 0 {
		return planets, apperror.ErrNotFound
	}

	return planets, nil
}

func (s *PlanetService) GetById(ctx context.Context, planetId string) (models.Planet, error) {
	planet, err := s.repo.Planet.GetById(ctx, planetId)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return planet, err
		}
		return planet, fmt.Errorf("failed to find planet by id. error: %w", err)
	}
	return planet, nil
}

func (s *PlanetService) Update(ctx context.Context, dto models.UpdatePlanetDTO) error {
	return nil
}

func (s *PlanetService) Delete(ctx context.Context, planetId string) error {
	err := s.repo.Planet.Delete(ctx, planetId)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return err
		}
		return fmt.Errorf("failed to delete planet. error: %w", err)
	}
	return nil
}
