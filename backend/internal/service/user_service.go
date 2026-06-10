package service

import (
	"github.com/absolute-achilles/plato/internal/repository"
)

type UserService interface {
	// GET
	// GetUser(ctx context.Context, id string) (*domain.User, error)
	// GetUsers(ctx context.Context) ([]domain.User, error)
	// GetStudent(ctx context.Context, id string) (*domain.Student, error)
	// GetTeacher(ctx context.Context, id string) (*domain.Teacher, error)
	// GetParent(ctx context.Context, id string) (*domain.Parent, error)
	// GetAdmin(ctx context.Context, id string) (*domain.Admin, error)
	//
	// // Create
	// CreateTeacher(ctx context.Context, req *dto.CreateTeacherRequest) (*domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// func (s *userService) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*domain.User, error) {
// 	user, err := s.repo.CreateUser(ctx, req)
// 	if err != nil {
// 		return nil, fmt.Errorf("userService.CreateUser: %w", err)
// 	}
// 	return user, nil
// }
//
// func (s *userService) GetUsers(ctx context.Context) ([]domain.User, error) {
// 	users, err := s.repo.GetAllUsers(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("userService.GetUsers: %w", err)
// 	}
// 	return users, nil
// }
//
// func (s *userService) GetUser(ctx context.Context, id string) (*domain.User, error) {
// 	user, err := s.repo.GetUserByID(ctx, id)
// 	if err != nil {
// 		return nil, fmt.Errorf("userService.GetUser: %w", err)
// 	}
// 	return user, nil
// }
// func (s *userService) GetStudent(ctx context.Context, id string) (*domain.Student, error) {
// 	student, err := s.repo.GetStudentByID(ctx, id)
// 	if err != nil {
// 		return nil, fmt.Errorf("userService.GetUser: %w", err)
// 	}
// 	return student, nil
// }
// func (s *userService) GetTeacher(ctx context.Context, id string) (*domain.Teacher, error) {
// 	teacher, err := s.repo.GetTeacherByID(ctx, id)
// 	if err != nil {
// 		return nil, fmt.Errorf("userService.GetUser: %w", err)
// 	}
// 	return teacher, nil
//
// }
// func (s *userService) GetParent(ctx context.Context, id string) (*domain.Parent, error) {
// 	guardian, err := s.repo.GetParentByID(ctx, id)
// 	if err != nil {
// 		return nil, fmt.Errorf("userService.GetUser: %w", err)
// 	}
// 	return guardian, nil
// }
// func (s *userService) GetAdmin(ctx context.Context, id string) (*domain.Admin, error) {
// 	admin, err := s.repo.GetAdminByID(ctx, id)
// 	if err != nil {
// 		return nil, fmt.Errorf("userService.GetUser: %w", err)
// 	}
// 	return admin, nil
// }
