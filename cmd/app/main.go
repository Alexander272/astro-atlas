package main

import (
	"os"

	"github.com/Alexander272/astro-atlas/internal/config"
	"github.com/Alexander272/astro-atlas/pkg/database/mongo"
	"github.com/Alexander272/astro-atlas/pkg/logger"
	"github.com/joho/godotenv"
)

// @title Astro Atlas
// @version 0.1
// @description API Server for Astro Atlas App

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logger.Init(os.Stdout)
	logger.Debug("init logger")
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}
	conf, err := config.Init("configs")
	if err != nil {
		logger.Fatalf("error initializing configs: %s", err.Error())
	}

	// Dependencies
	mongoClient, err := mongo.NewClient(conf.Mongo.URI, conf.Mongo.User, conf.Mongo.Password)
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}
	db := mongoClient.Database(conf.Mongo.Name)

	logger.Debug(db)

	// client, err := redis.NewRedisClient(redis.Config{
	// 	Host:     conf.Redis.Host,
	// 	Port:     conf.Redis.Port,
	// 	DB:       conf.Redis.DB,
	// 	Password: conf.Redis.Password,
	// })
	// if err != nil {
	// 	logger.Fatalf("failed to initialize redis %s", err.Error())
	// }

	// hasher := hasher.NewBcryptHasher(conf.Auth.Bcrypt.MinCost, conf.Auth.Bcrypt.DefaultCost, conf.Auth.Bcrypt.MaxCost)
	// tokenManager, err := auth.NewManager(conf.Auth.JWT.Key)
	// if err != nil {
	// 	logger.Fatalf("failed to initialize token manager: %s", err.Error())
	// }

	// Services, Repos & API Handlers
}
