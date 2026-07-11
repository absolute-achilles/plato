package service

import (
	"context"
	"testing"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/dto"
	"github.com/absolute-achilles/plato/internal/repository"
	"github.com/stretchr/testify/require"
)

type mockTeacherRepository struct {
	created []*domain.Teacher
}

func (m *mockTeacherRepository) Create(ctx context.Context, teacher *domain.Teacher) error {
	teacher.ID = "teacher-" + string(teacher.Email)
	m.created = append(m.created, teacher)
	return nil
}

type mockStudentRepository struct {
	created []*domain.Student
}

func (m *mockStudentRepository) Create(ctx context.Context, student *domain.Student) error {
	student.ID = "student-" + string(student.Email)
	m.created = append(m.created, student)
	return nil
}

type mockParentRepository struct {
	created []*domain.Parent
}

func (m *mockParentRepository) Create(ctx context.Context, parent *domain.Parent) error {
	parent.ID = "parent-" + string(parent.Email)
	m.created = append(m.created, parent)
	return nil
}

func TestAdminService_CreateTeacher(t *testing.T) {
	ctx := context.Background()
	teacherRepo := &mockTeacherRepository{}
	studentRepo := &mockStudentRepository{}
	parentRepo := &mockParentRepository{}

	svc := NewAdminService(teacherRepo, studentRepo, parentRepo)

	resp, err := svc.CreateTeacher(ctx, &dto.CreateTeacherRequest{
		Username:   "budi.santoso",
		Email:      "budi@plato.edu",
		Password:   "password123",
		Department: domain.DepartmentMathematics,
	})
	require.NoError(t, err)
	require.Equal(t, "budi.santoso", resp.Username)
	require.Equal(t, domain.DepartmentMathematics, resp.Department)
	require.Equal(t, domain.RoleTeacher, resp.Role)
	require.Len(t, teacherRepo.created, 1)

	// The real repository hashes the password; the mock just stores the plaintext value passed by the service.
	require.Equal(t, "password123", teacherRepo.created[0].HashPassword)
}

func TestAdminService_CreateStudent(t *testing.T) {
	ctx := context.Background()
	teacherRepo := &mockTeacherRepository{}
	studentRepo := &mockStudentRepository{}
	parentRepo := &mockParentRepository{}

	svc := NewAdminService(teacherRepo, studentRepo, parentRepo)

	resp, err := svc.CreateStudent(ctx, &dto.CreateStudentRequest{
		Username:   "andi.wijaya",
		Email:      "andi@plato.edu",
		Password:   "password123",
		GradeLevel: domain.GradeLevel10,
	})
	require.NoError(t, err)
	require.Equal(t, domain.GradeLevel10, resp.GradeLevel)
	require.Len(t, studentRepo.created, 1)
}

func TestAdminService_CreateParent(t *testing.T) {
	ctx := context.Background()
	teacherRepo := &mockTeacherRepository{}
	studentRepo := &mockStudentRepository{}
	parentRepo := &mockParentRepository{}

	svc := NewAdminService(teacherRepo, studentRepo, parentRepo)

	resp, err := svc.CreateParent(ctx, &dto.CreateParentRequest{
		Username:   "siti.aminah",
		Email:      "siti@plato.edu",
		Password:   "password123",
		Type:       domain.ParentRelationshipMother,
		StudentIDs: []string{"student-1"},
	})
	require.NoError(t, err)
	require.Equal(t, domain.ParentRelationshipMother, resp.Type)
	require.Equal(t, []string{"student-1"}, resp.StudentIDs)
	require.Len(t, parentRepo.created, 1)
}

func TestAdminService_CreateTeacherDuplicate(t *testing.T) {
	ctx := context.Background()
	studentRepo := &mockStudentRepository{}
	parentRepo := &mockParentRepository{}

	// Override the mock to simulate a duplicate error
	dupeTeacherRepo := &duplicateTeacherRepository{}
	svc := NewAdminService(dupeTeacherRepo, studentRepo, parentRepo)

	_, err := svc.CreateTeacher(ctx, &dto.CreateTeacherRequest{
		Username:   "budi.santoso",
		Email:      "budi@plato.edu",
		Password:   "password123",
		Department: domain.DepartmentMathematics,
	})
	require.ErrorIs(t, err, domain.ErrDuplicate)
}

type duplicateTeacherRepository struct{}

func (d *duplicateTeacherRepository) Create(ctx context.Context, teacher *domain.Teacher) error {
	return repository.ErrDuplicateEmail
}
