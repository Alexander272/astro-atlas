package repository

import (
	planetRepo "github.com/Alexander272/astro-atlas/internal/planet/repository"
	planetDb "github.com/Alexander272/astro-atlas/internal/planet/repository/mongodb"
	userRepo "github.com/Alexander272/astro-atlas/internal/user/repository"
	userDb "github.com/Alexander272/astro-atlas/internal/user/repository/mongodb"
	sessionDb "github.com/Alexander272/astro-atlas/internal/user/repository/redis"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	Session userRepo.ISession
	User    userRepo.IUser
	Planet  planetRepo.IPlanet
	System  planetRepo.ISystem
}

func NewRepositories(db *mongo.Database, client redis.Cmdable) *Repositories {
	return &Repositories{
		Session: sessionDb.NewSessionRepo(client),
		User:    userDb.NewUserRepo(db, usersCollection),
		Planet:  planetDb.NewPlanetRepo(db, planetCollection),
		System:  planetDb.NewSystemRepo(db, systemsCollection),
	}
}
