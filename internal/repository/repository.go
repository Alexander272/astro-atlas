package repository

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct{}

func NewRepositories(db *mongo.Database, client *redis.Client) *Repositories {
	return &Repositories{}
}
