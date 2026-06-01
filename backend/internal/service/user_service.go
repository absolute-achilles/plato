package service

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/dto"
	"github.com/absolute-achilles/plato/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*domain.User, error)
	GetUser(ctx context.Context, id int64) (*domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*domain.User, error) {
	user, err := s.repo.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("userService.CreateUser: %w", err)
	}
	return user, nil
}

func (s *userService) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("userService.GetUser: %w", err)
	}
	return user, nil
}
