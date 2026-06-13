package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CourseRepository interface {
	Get(ctx context.Context, courseID string) (*domain.Course, error)
	GetTeacherCoursesByID(ctx context.Context, teacherID string) ([]*domain.Course, error)
	GetTeacherCoursesByEmail(ctx context.Context, teacherEmail string) ([]*domain.Course, error)
	Create(ctx context.Context, course *domain.Course) (id string, err error)
	Delete(ctx context.Context, courseID string) error
}

type courseRepository struct {
	db *pgxpool.Pool
}

func NewCourseRepository(db *pgxpool.Pool) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) Get(ctx context.Context, courseID string) (*domain.Course, error) {
	query := `
	SELECT id, teacher_id, name, description, created_at, updated_at 
	FROM courses WHERE id = $1`

	rows, err := r.db.Query(ctx, query, courseID)
	if err != nil {
		return nil, fmt.Errorf("courseRepository.Get: %w", err)
	}
	course, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.Course])
	if err != nil {
		return nil, fmt.Errorf("courseRepository.Get: %w", err)
	}

	return course, nil
}

func (r *courseRepository) Create(ctx context.Context, course *domain.Course) (id string, err error) {
	if course == nil {
		return "", fmt.Errorf("courseRepository.Create: Empty Course Request")
	}

	if course.TeacherID == "" {
		return "", fmt.Errorf("courseRepository.Create: Empty Teacher ID")
	}

	insertQuery := `
		INSERT INTO courses (teacher_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	if err := r.db.QueryRow(ctx, insertQuery, course.TeacherID, course.Name, course.Description).Scan(&id); err != nil {
		return "", fmt.Errorf("courseRepository.Create: %w", err)
	}

	return id, nil
}

func (r *courseRepository) Delete(ctx context.Context, courseID string) error {
	query := `DELETE FROM courses WHERE id = $1`
	result, err := r.db.Exec(ctx, query, courseID)
	if err != nil {
		return fmt.Errorf("courseRepository.Delete: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("courseRepository.Delete: course not found with ID: %s", courseID)
	}

	return nil
}

func (r *courseRepository) GetTeacherCoursesByID(ctx context.Context, teacherID string) ([]*domain.Course, error) {
	if teacherID == "" {
		return nil, fmt.Errorf("courseRepository.GetTeacherCoursesByID: Empty Teacher ID")
	}

	query := `
	SELECT id, teacher_id, name, description, created_at, updated_at 
	FROM courses WHERE teacher_id = $1`

	rows, err := r.db.Query(ctx, query, teacherID)
	if err != nil {
		return nil, fmt.Errorf("courseRepository.GetTeacherCoursesByID: %w", err)
	}

	courses, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.Course])
	if err != nil {
		return nil, fmt.Errorf("courseRepository.GetTeacherCoursesByID: %w", err)
	}

	return courses, nil
}

func (r *courseRepository) GetTeacherCoursesByEmail(ctx context.Context, teacherEmail string) ([]*domain.Course, error) {
	if teacherEmail == "" {
		return nil, fmt.Errorf("courseRepository.GetTeacherCoursesByID: Empty Teacher ID")
	}

	query := `SELECT id FROM users WHERE email = $1`

	var teacherID string
	if err := r.db.QueryRow(ctx, query, teacherEmail).Scan(&teacherID); err != nil {
		return nil, fmt.Errorf("courseRepository.GetTeacherCoursesByEmail: %w", err)
	}

	return r.GetTeacherCoursesByID(ctx, teacherID)
}
