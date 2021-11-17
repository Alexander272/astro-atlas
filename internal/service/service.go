package service

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/Alexander272/astro-atlas/internal/models"
	planetService "github.com/Alexander272/astro-atlas/internal/planet/service"
	"github.com/Alexander272/astro-atlas/internal/repository"
	userService "github.com/Alexander272/astro-atlas/internal/user/service"
	"github.com/Alexander272/astro-atlas/pkg/auth"
	"github.com/Alexander272/astro-atlas/pkg/hasher"
	"github.com/Alexander272/astro-atlas/pkg/storage"
)

type IFile interface {
	Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader, path, filename string) (models.File, error)
	Delete(ctx context.Context, path, filename string) error
}

type Services struct {
	Auth   userService.IAuth
	User   userService.IUser
	System planetService.ISystems
	Planet planetService.IPlanet
	File   IFile
}

type Deps struct {
	Repos           *repository.Repositories
	StorageProvider storage.Provider
	Hasher          hasher.IPasswordHasher
	TokenManager    auth.ITokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	Domain          string
}

func NewServices(deps Deps) *Services {
	return &Services{
		Auth: userService.NewAuthService(deps.Repos.User, deps.Repos.Session, deps.TokenManager, deps.Hasher, deps.AccessTokenTTL,
			deps.RefreshTokenTTL, deps.Domain),
		User:   userService.NewUserService(deps.Repos.User, deps.Hasher),
		System: planetService.NewSystemService(deps.Repos.System),
		Planet: planetService.NewPlanetService(deps.Repos.Planet),
		File:   NewFileService(deps.StorageProvider),
	}
}
