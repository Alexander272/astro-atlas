package service

import (
	"time"

	"github.com/Alexander272/astro-atlas/internal/repository"
	"github.com/Alexander272/astro-atlas/pkg/auth"
	"github.com/Alexander272/astro-atlas/pkg/hasher"
)

type Services struct{}

type Deps struct {
	Repos *repository.Repositories
	// StorageProvider        storage.Provider
	Hasher          hasher.IPasswordHasher
	TokenManager    auth.ITokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	Domain          string
}

func NewServices(deps Deps) *Services {
	return &Services{}
}
