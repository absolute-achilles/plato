package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AttendanceRepository interface {
	Create(ctx context.Context, attendance *domain.Attendance) (id string, err error)
	Get(ctx context.Context, id string) (*domain.Attendance, error)
	Update(ctx context.Context, attendance *domain.Attendance) error
	GetByStudentAndModule(ctx context.Context, studentID, moduleID string) (*domain.Attendance, error)
	GetModuleAttendances(ctx context.Context, moduleID string) ([]*domain.Attendance, error)
	GetStudentAttendances(ctx context.Context, studentID string) ([]*domain.Attendance, error)
	Delete(ctx context.Context, id string) error
}

type attendanceRepository struct {
	db *pgxpool.Pool
}

func NewAttendanceRepository(db *pgxpool.Pool) AttendanceRepository {
	return &attendanceRepository{db: db}
}

func (r *attendanceRepository) Create(ctx context.Context, attendance *domain.Attendance) (id string, err error) {
	if attendance == nil {
		return "", fmt.Errorf("attendanceRepository.Create: Empty Attendance Request")
	}

	insertQuery := `
		INSERT INTO attendances (student_id, module_id, status, notes)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	if err := r.db.QueryRow(
		ctx,
		insertQuery,
		attendance.StudentID,
		attendance.ModuleID,
		attendance.Status,
		attendance.Notes,
	).Scan(&id); err != nil {
		return "", fmt.Errorf("attendanceRepository.Create: %w", err)
	}

	return id, nil
}

func (r *attendanceRepository) Get(ctx context.Context, id string) (*domain.Attendance, error) {
	query := `
	SELECT id, student_id, module_id, status, recorded_at, notes 
	FROM attendances WHERE id = $1`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("attendanceRepository.Get: %w", err)
	}

	attendance, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.Attendance])
	if err != nil {
		return nil, fmt.Errorf("attendanceRepository.Get: %w", err)
	}

	return attendance, nil
}

func (r *attendanceRepository) Update(ctx context.Context, attendance *domain.Attendance) error {
	if attendance == nil || attendance.ID == "" {
		return fmt.Errorf("attendanceRepository.Update: Invalid Attendance Request")
	}

	query := `
		UPDATE attendances 
		SET status = $1, notes = $2, recorded_at = now() 
		WHERE id = $3
	`

	result, err := r.db.Exec(ctx, query, attendance.Status, attendance.Notes, attendance.ID)
	if err != nil {
		return fmt.Errorf("attendanceRepository.Update: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("attendanceRepository.Update: attendance not found with ID: %s", attendance.ID)
	}

	return nil
}

func (r *attendanceRepository) GetByStudentAndModule(ctx context.Context, studentID, moduleID string) (*domain.Attendance, error) {
	query := `
	SELECT id, student_id, module_id, status, recorded_at, notes 
	FROM attendances WHERE student_id = $1 AND module_id = $2`

	rows, err := r.db.Query(ctx, query, studentID, moduleID)
	if err != nil {
		return nil, fmt.Errorf("attendanceRepository.GetByStudentAndModule: %w", err)
	}

	attendance, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.Attendance])
	if err != nil {
		return nil, fmt.Errorf("attendanceRepository.GetByStudentAndModule: %w", err)
	}

	return attendance, nil
}

func (r *attendanceRepository) GetModuleAttendances(ctx context.Context, moduleID string) ([]*domain.Attendance, error) {
	query := `
	SELECT id, student_id, module_id, status, recorded_at, notes 
	FROM attendances WHERE module_id = $1`

	rows, err := r.db.Query(ctx, query, moduleID)
	if err != nil {
		return nil, fmt.Errorf("attendanceRepository.GetModuleAttendances: %w", err)
	}

	attendances, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.Attendance])
	if err != nil {
		return nil, fmt.Errorf("attendanceRepository.GetModuleAttendances: %w", err)
	}

	return attendances, nil
}

func (r *attendanceRepository) GetStudentAttendances(ctx context.Context, studentID string) ([]*domain.Attendance, error) {
	query := `
	SELECT id, student_id, module_id, status, recorded_at, notes 
	FROM attendances WHERE student_id = $1`

	rows, err := r.db.Query(ctx, query, studentID)
	if err != nil {
		return nil, fmt.Errorf("attendanceRepository.GetStudentAttendances: %w", err)
	}

	attendances, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.Attendance])
	if err != nil {
		return nil, fmt.Errorf("attendanceRepository.GetStudentAttendances: %w", err)
	}

	return attendances, nil
}

func (r *attendanceRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM attendances WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("attendanceRepository.Delete: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("attendanceRepository.Delete: attendance not found with ID: %s", id)
	}

	return nil
}
