package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EnrollmentRepository interface {
	Create(ctx context.Context, enrollment *domain.Enrollment) (id string, err error)
	Get(ctx context.Context, id string) (*domain.Enrollment, error)
	GetByStudentAndCourse(ctx context.Context, studentID, courseID string) (*domain.Enrollment, error)
	GetStudentEnrollments(ctx context.Context, studentID string) ([]*domain.Enrollment, error)
	GetCourseEnrollments(ctx context.Context, courseID string) ([]*domain.Enrollment, error)
	Delete(ctx context.Context, id string) error
}

type enrollmentRepository struct {
	db *pgxpool.Pool
}

func NewEnrollmentRepository(db *pgxpool.Pool) EnrollmentRepository {
	return &enrollmentRepository{db: db}
}

func (r *enrollmentRepository) Create(ctx context.Context, enrollment *domain.Enrollment) (id string, err error) {
	if enrollment == nil {
		return "", fmt.Errorf("enrollmentRepository.Create: Empty Enrollment Request")
	}

	insertQuery := `
		INSERT INTO enrollments (student_id, course_id)
		VALUES ($1, $2)
		RETURNING id
	`

	if err := r.db.QueryRow(
		ctx,
		insertQuery,
		enrollment.StudentID,
		enrollment.CourseID,
	).Scan(&id); err != nil {
		return "", fmt.Errorf("enrollmentRepository.Create: %w", err)
	}

	return id, nil
}

func (r *enrollmentRepository) Get(ctx context.Context, id string) (*domain.Enrollment, error) {
	query := `
	SELECT id, student_id, course_id, enrolled_at 
	FROM enrollments WHERE id = $1`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("enrollmentRepository.Get: %w", err)
	}

	enrollment, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.Enrollment])
	if err != nil {
		return nil, fmt.Errorf("enrollmentRepository.Get: %w", err)
	}

	return enrollment, nil
}

func (r *enrollmentRepository) GetByStudentAndCourse(ctx context.Context, studentID, courseID string) (*domain.Enrollment, error) {
	query := `
	SELECT id, student_id, course_id, enrolled_at 
	FROM enrollments WHERE student_id = $1 AND course_id = $2`

	rows, err := r.db.Query(ctx, query, studentID, courseID)
	if err != nil {
		return nil, fmt.Errorf("enrollmentRepository.GetByStudentAndCourse: %w", err)
	}

	enrollment, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.Enrollment])
	if err != nil {
		return nil, fmt.Errorf("enrollmentRepository.GetByStudentAndCourse: %w", err)
	}

	return enrollment, nil
}

func (r *enrollmentRepository) GetStudentEnrollments(ctx context.Context, studentID string) ([]*domain.Enrollment, error) {
	query := `
	SELECT id, student_id, course_id, enrolled_at 
	FROM enrollments WHERE student_id = $1`

	rows, err := r.db.Query(ctx, query, studentID)
	if err != nil {
		return nil, fmt.Errorf("enrollmentRepository.GetStudentEnrollments: %w", err)
	}

	enrollments, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.Enrollment])
	if err != nil {
		return nil, fmt.Errorf("enrollmentRepository.GetStudentEnrollments: %w", err)
	}

	return enrollments, nil
}

func (r *enrollmentRepository) GetCourseEnrollments(ctx context.Context, courseID string) ([]*domain.Enrollment, error) {
	query := `
	SELECT id, student_id, course_id, enrolled_at 
	FROM enrollments WHERE course_id = $1`

	rows, err := r.db.Query(ctx, query, courseID)
	if err != nil {
		return nil, fmt.Errorf("enrollmentRepository.GetCourseEnrollments: %w", err)
	}

	enrollments, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.Enrollment])
	if err != nil {
		return nil, fmt.Errorf("enrollmentRepository.GetCourseEnrollments: %w", err)
	}

	return enrollments, nil
}

func (r *enrollmentRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM enrollments WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("enrollmentRepository.Delete: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("enrollmentRepository.Delete: enrollment not found with ID: %s", id)
	}

	return nil
}
