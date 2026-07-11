package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/dto"
	"github.com/absolute-achilles/plato/internal/repository"
)

type AdminService interface {
	CreateTeacher(ctx context.Context, req *dto.CreateTeacherRequest) (*dto.TeacherResponse, error)
	CreateStudent(ctx context.Context, req *dto.CreateStudentRequest) (*dto.StudentResponse, error)
	CreateParent(ctx context.Context, req *dto.CreateParentRequest) (*dto.ParentResponse, error)
}

type adminService struct {
	teacherRepo repository.TeacherRepository
	studentRepo repository.StudentRepository
	parentRepo  repository.ParentRepository
}

func NewAdminService(
	teacherRepo repository.TeacherRepository,
	studentRepo repository.StudentRepository,
	parentRepo repository.ParentRepository,
) AdminService {
	return &adminService{
		teacherRepo: teacherRepo,
		studentRepo: studentRepo,
		parentRepo:  parentRepo,
	}
}

func stringPtrOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func (s *adminService) CreateTeacher(ctx context.Context, req *dto.CreateTeacherRequest) (*dto.TeacherResponse, error) {
	teacher := &domain.Teacher{
		User: domain.User{
			Username:     req.Username,
			Email:        req.Email,
			HashPassword: req.Password,
			PhoneNumber:  stringPtrOrNil(req.PhoneNumber),
		},
		Department: req.Department,
	}

	if err := s.teacherRepo.Create(ctx, teacher); err != nil {
		if errors.Is(err, repository.ErrDuplicateEmail) || errors.Is(err, repository.ErrDuplicateName) || errors.Is(err, domain.ErrDuplicate) {
			return nil, domain.ErrDuplicate
		}
		return nil, fmt.Errorf("adminService.CreateTeacher: %w", err)
	}

	return &dto.TeacherResponse{
		ID:          teacher.ID,
		Username:    teacher.Username,
		Name:        teacher.Username,
		Email:       teacher.Email,
		Role:        domain.RoleTeacher,
		PhoneNumber: teacher.PhoneNumber,
		Department:  teacher.Department,
		CreatedAt:   teacher.CreatedAt,
	}, nil
}

func (s *adminService) CreateStudent(ctx context.Context, req *dto.CreateStudentRequest) (*dto.StudentResponse, error) {
	student := &domain.Student{
		User: domain.User{
			Username:     req.Username,
			Email:        req.Email,
			HashPassword: req.Password,
			PhoneNumber:  stringPtrOrNil(req.PhoneNumber),
		},
		GradeLevel: req.GradeLevel,
	}

	if err := s.studentRepo.Create(ctx, student); err != nil {
		if errors.Is(err, repository.ErrDuplicateEmail) || errors.Is(err, repository.ErrDuplicateName) || errors.Is(err, domain.ErrDuplicate) {
			return nil, domain.ErrDuplicate
		}
		return nil, fmt.Errorf("adminService.CreateStudent: %w", err)
	}

	return &dto.StudentResponse{
		ID:          student.ID,
		Username:    student.Username,
		Name:        student.Username,
		Email:       student.Email,
		Role:        domain.RoleStudent,
		PhoneNumber: student.PhoneNumber,
		GradeLevel:  student.GradeLevel,
		CreatedAt:   student.CreatedAt,
	}, nil
}

func (s *adminService) CreateParent(ctx context.Context, req *dto.CreateParentRequest) (*dto.ParentResponse, error) {
	parent := &domain.Parent{
		User: domain.User{
			Username:     req.Username,
			Email:        req.Email,
			HashPassword: req.Password,
			PhoneNumber:  stringPtrOrNil(req.PhoneNumber),
		},
		Type:       req.Type,
		StudentIDs: req.StudentIDs,
	}

	if err := s.parentRepo.Create(ctx, parent); err != nil {
		if errors.Is(err, repository.ErrDuplicateEmail) || errors.Is(err, repository.ErrDuplicateName) || errors.Is(err, domain.ErrDuplicate) {
			return nil, domain.ErrDuplicate
		}
		return nil, fmt.Errorf("adminService.CreateParent: %w", err)
	}

	return &dto.ParentResponse{
		ID:          parent.ID,
		Username:    parent.Username,
		Name:        parent.Username,
		Email:       parent.Email,
		Role:        domain.RoleParent,
		PhoneNumber: parent.PhoneNumber,
		Type:        parent.Type,
		StudentIDs:  parent.StudentIDs,
		CreatedAt:   parent.CreatedAt,
	}, nil
}
