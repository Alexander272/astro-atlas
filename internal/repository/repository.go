package repository

import (
	planetRepo "github.com/Alexander272/astro-atlas/internal/planet/repository"
	planetDb "github.com/Alexander272/astro-atlas/internal/planet/repository/mongodb"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	Planet planetRepo.IPlanet
	System planetRepo.ISystem
}

func NewRepositories(db *mongo.Database, client *redis.Client) *Repositories {
	return &Repositories{
		Planet: planetDb.NewPlanetRepo(db, planetCollection),
		System: planetDb.NewSystemRepo(db, systemsCollection),
	}
}
