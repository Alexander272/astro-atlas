package service

import (
	"context"

	"github.com/Alexander272/astro-atlas/internal/planet/models"
)

type IPlanet interface {
	Create(ctx context.Context, dto models.CreatePlanetDTO) (string, error)
	GetList(ctx context.Context, systemId string) ([]models.PlanetShort, error)
	GetById(ctx context.Context, planetId string) (models.Planet, error)
	Update(ctx context.Context, dto models.UpdatePlanetDTO) error
	Delete(ctx context.Context, planetId string) error
}

type ISystems interface {
	Create(ctx context.Context, dto models.CreateSystemDTO) (string, error)
	GetList(ctx context.Context) ([]models.SystemShort, error)
	GetById(ctx context.Context, systemId string) (models.System, error)
	Update(ctx context.Context, dto models.UpdateSystemDTO) error
	Delete(ctx context.Context, systemId string) error
}
