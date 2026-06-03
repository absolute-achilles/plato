package service

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/dto"
	"github.com/absolute-achilles/plato/internal/repository"
)

type UserService interface {
	// GET
	GetUser(ctx context.Context, id string, role domain.Role) (*domain.User, error)
	GetUsers(ctx context.Context, role domain.Role) ([]domain.User, error)
	GetStudent(ctx context.Context, id string) (*domain.Student, error)
	GetTeacher(ctx context.Context, id string) (*domain.Teacher, error)
	GetGuardian(ctx context.Context, id string) (*domain.Guardian, error)
	GetAdmin(ctx context.Context, id string) (*domain.Admin, error)

	// Create
	CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*domain.User, error)
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

func (s *userService) GetUsers(ctx context.Context, role domain.Role) ([]domain.User, error) {
	users, err := s.repo.GetAllUsers(ctx, role)
	if err != nil {
		return nil, fmt.Errorf("userService.GetUsers: %w", err)
	}
	return users, nil
}

func (s *userService) GetUser(ctx context.Context, id string, role domain.Role) (*domain.User, error) {
	user, err := s.repo.GetUserByID(ctx, id, role)
	if err != nil {
		return nil, fmt.Errorf("userService.GetUser: %w", err)
	}
	return user, nil
}
func (s *userService) GetStudent(ctx context.Context, id string) (*domain.Student, error) {
	student, err := s.repo.GetStudentByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("userService.GetUser: %w", err)
	}
	return student, nil
}
func (s *userService) GetTeacher(ctx context.Context, id string) (*domain.Teacher, error) {
	teacher, err := s.repo.GetTeacherByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("userService.GetUser: %w", err)
	}
	return teacher, nil

}
func (s *userService) GetGuardian(ctx context.Context, id string) (*domain.Guardian, error) {
	guardian, err := s.repo.GetGuardianByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("userService.GetUser: %w", err)
	}
	return guardian, nil
}
func (s *userService) GetAdmin(ctx context.Context, id string) (*domain.Admin, error) {
	admin, err := s.repo.GetAdminByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("userService.GetUser: %w", err)
	}
	return admin, nil
}
