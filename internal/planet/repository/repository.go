package repository

import (
	"context"

	"github.com/Alexander272/astro-atlas/internal/planet/models"
)

type IPlanet interface {
	Create(ctx context.Context, planet models.Planet) (string, error)
	GetList(ctx context.Context, systemId string) ([]models.PlanetShort, error)
	GetById(ctx context.Context, planetId string) (models.Planet, error)
	Update(ctx context.Context, planet models.Planet) error
	Delete(ctx context.Context, planetId string) error
}

type ISystem interface {
	Create(ctx context.Context, system models.System) (string, error)
	GetList(ctx context.Context) ([]models.SystemShort, error)
	GetById(ctx context.Context, id string) (models.System, error)
	Update(ctx context.Context, system models.System) error
	Delete(ctx context.Context, id string) error
}
