package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alexander272/astro-atlas/internal/planet/models"
	"github.com/Alexander272/astro-atlas/internal/planet/repository"
	"github.com/Alexander272/astro-atlas/pkg/apperror"
)

type PlanetService struct {
	repo repository.IPlanet
}

func NewPlanetService(repo repository.IPlanet) *PlanetService {
	return &PlanetService{
		repo: repo,
	}
}

func (s *PlanetService) Create(ctx context.Context, dto models.CreatePlanetDTO) (planetId string, err error) {
	planet := models.NewPlanet(dto)
	planetId, err = s.repo.Create(ctx, planet)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return planetId, err
		}
		return planetId, fmt.Errorf("failed to create planet. error: %w", err)
	}

	return planetId, nil
}

func (s *PlanetService) GetList(ctx context.Context, systemId string) (planets []models.PlanetShort, err error) {
	planets, err = s.repo.GetList(ctx, systemId)
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

func (s *PlanetService) GetById(ctx context.Context, planetId string) (planet models.Planet, err error) {
	planet, err = s.repo.GetById(ctx, planetId)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return planet, err
		}
		return planet, fmt.Errorf("failed to find planet by id. error: %w", err)
	}
	return planet, nil
}

func (s *PlanetService) Update(ctx context.Context, dto models.UpdatePlanetDTO) error {
	err := s.repo.Update(ctx, models.Planet(dto))
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return err
		}
		return fmt.Errorf("failed to update planet. error: %w", err)
	}
	return nil
}

func (s *PlanetService) Delete(ctx context.Context, planetId string) error {
	err := s.repo.Delete(ctx, planetId)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return err
		}
		return fmt.Errorf("failed to delete planet. error: %w", err)
	}
	return nil
}
