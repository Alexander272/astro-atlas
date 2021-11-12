package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Alexander272/astro-atlas/internal/user/models"
	"github.com/Alexander272/astro-atlas/internal/user/repository"
	"github.com/Alexander272/astro-atlas/pkg/apperror"
	"github.com/Alexander272/astro-atlas/pkg/hasher"
	"github.com/Alexander272/astro-atlas/pkg/logger"
)

type UserService struct {
	repo   repository.IUser
	hasher hasher.IPasswordHasher
}

func NewUserService(repo repository.IUser, hasher hasher.IPasswordHasher) *UserService {
	return &UserService{
		repo:   repo,
		hasher: hasher,
	}
}

func (s *UserService) Create(ctx context.Context, dto models.CreateUserDTO) (id string, err error) {
	candidate, err := s.repo.GetByEmail(ctx, dto.Email)
	if err != nil {
		if !errors.Is(err, apperror.ErrNotFound) {
			return id, fmt.Errorf("failed to get user by email. error: %w", err)
		}
	}
	if (candidate != models.User{}) {
		return id, fmt.Errorf("user with the same email address already exists")
	}

	user := models.NewUser(dto)
	pasHah, err := s.hasher.HashPassword(user.Password)
	if err != nil {
		logger.Errorf("failed to create user due to error %v", err)
		return id, fmt.Errorf("failed to create user. error: %w", err)
	}

	user.Password = pasHah
	id, err = s.repo.Create(ctx, user)
	if err != nil {
		return id, fmt.Errorf("failed to create user. error: %w", err)
	}

	return id, nil
}

func (s *UserService) GetAll(ctx context.Context) (users []models.User, err error) {
	users, err = s.repo.GetAll(ctx)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return users, err
		}
		return users, fmt.Errorf("failed to get users. error: %w", err)
	}
	if len(users) == 0 {
		return users, apperror.ErrNotFound
	}

	return users, nil
}

func (s *UserService) GetById(ctx context.Context, userId string) (u models.User, err error) {
	u, err = s.repo.GetById(ctx, userId)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return u, err
		}
		return u, fmt.Errorf("failed to get user by id. error: %w", err)
	}

	return u, nil
}

func (s *UserService) Update(ctx context.Context, dto models.UpdateUserDTO) error {
	updateUser := models.UpdateUser(dto)
	err := s.repo.Update(ctx, updateUser)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return err
		}
		return fmt.Errorf("failed to update user. error: %w", err)
	}
	return nil
}

func (s *UserService) Delete(ctx context.Context, userId string) error {
	err := s.repo.Delete(ctx, userId)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return err
		}
		return fmt.Errorf("failed to delete user. error: %w", err)
	}
	return nil
}
