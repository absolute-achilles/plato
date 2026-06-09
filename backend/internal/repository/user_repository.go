package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/dto"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	// if role is empty, will query from user table instead
	GetUserByID(ctx context.Context, id string, role domain.Role) (*domain.User, error)
	GetAllUsers(ctx context.Context, role domain.Role) ([]domain.User, error)

	GetStudentByID(ctx context.Context, id string) (*domain.Student, error)
	GetTeacherByID(ctx context.Context, id string) (*domain.Teacher, error)
	GetParentByID(ctx context.Context, id string) (*domain.Parent, error)
	GetAdminByID(ctx context.Context, id string) (*domain.Admin, error)

	Create(ctx context.Context, req *dto.CreateUserRequest) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func queryUserByID[T any](ctx context.Context, db *sqlx.DB, id string, role domain.Role) (*T, error) {
	if role != "" && role != domain.RoleAdmin && role != domain.RoleParent && role != domain.RoleStudent && role != domain.RoleTeacher {
		return nil, fmt.Errorf("invalid role: %s", role)
	}

	var user T
	var query string
	if role == "" {
		// Fallback: Query just the base user table if no specific role is provided
		query = `SELECT id, name, email, role FROM "user" WHERE id = $1`
	} else {
		query = fmt.Sprintf(`
            SELECT u.id, u.name, u.email, u.role 
            FROM "user" u 
            INNER JOIN %s r ON u.id = r.user_id 
            WHERE u.id = $1
        `, role)
	}
	err := db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByID: %w", err)
	}
	return &user, nil
}

func queryUsers[T any](ctx context.Context, db *sqlx.DB, role domain.Role) ([]T, error) {
	if role != "" && role != domain.RoleAdmin && role != domain.RoleParent && role != domain.RoleStudent && role != domain.RoleTeacher {
		return nil, fmt.Errorf("invalid role: %s", role)
	}

	users := make([]T, 0, 100)
	var query string
	if role == "" {
		// Fallback: Query just the base user table if no specific role is provided
		query = `SELECT id, name, email, role FROM "user"`
	} else {
		query = fmt.Sprintf(`
            SELECT u.id, u.name, u.email, u.role 
            FROM "user" u 
            INNER JOIN %s r ON u.id = r.user_id 
        `, role)
	}
	err := db.GetContext(ctx, &users, query)
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByID: %w", err)
	}
	return users, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id string, role domain.Role) (*domain.User, error) {
	user, err := queryUserByID[domain.User](ctx, r.db, id, role)
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByID: %w", err)
	}
	return user, nil
}

func (r *userRepository) GetAllUsers(ctx context.Context, role domain.Role) ([]domain.User, error) {
	if role != "" && role != domain.RoleAdmin && role != domain.RoleParent && role != domain.RoleStudent && role != domain.RoleTeacher {
		return nil, fmt.Errorf("invalid role: %s", role)
	}

	users, err := queryUsers[domain.User](ctx, r.db, role)
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetAllUsers: %w", err)
	}

	return users, nil
}

func (r *userRepository) GetStudentByID(ctx context.Context, id string) (*domain.Student, error) {
	student, err := queryUserByID[domain.Student](ctx, r.db, id, domain.RoleStudent)
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByID: %w", err)
	}
	return student, nil
}

func (r *userRepository) GetTeacherByID(ctx context.Context, id string) (*domain.Teacher, error) {
	teacher, err := queryUserByID[domain.Teacher](ctx, r.db, id, domain.RoleTeacher)
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByID: %w", err)
	}
	return teacher, nil
}

func (r *userRepository) GetParentByID(ctx context.Context, id string) (*domain.Parent, error) {
	guardian, err := queryUserByID[domain.Parent](ctx, r.db, id, domain.RoleParent)
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByID: %w", err)
	}
	return guardian, nil
}

func (r *userRepository) GetAdminByID(ctx context.Context, id string) (*domain.Admin, error) {
	admin, err := queryUserByID[domain.Admin](ctx, r.db, id, domain.RoleAdmin)
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByID: %w", err)
	}
	return admin, nil
}

func (r *userRepository) Create(ctx context.Context, req *dto.CreateUserRequest) (*domain.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	var user domain.User
	insertBaseQuery := `
        INSERT INTO "user" (name, email, password, role)
        VALUES ($1, $2, $3, UPPER($4))
        RETURNING id, name, email, role`
	err = r.db.QueryRowxContext(ctx, insertBaseQuery, req.Name, req.Email, req.Role).StructScan(&user)
	if err != nil {
		return nil, fmt.Errorf("userRepository.Create - base user: %w", err)
	}

	insertSubclassQuery := fmt.Sprintf(`INSERT INTO %s (user_id) VALUES ($1)`, strings.ToLower(string(user.Role)))
	_, err = tx.ExecContext(ctx, insertSubclassQuery, user.ID)
	if err != nil {
		return nil, fmt.Errorf("userRepository.Create - subclass (%s): %w", user.Role, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &user, nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM "user" WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("userRepository.Delete: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user not found with ID: %s", id)
	}

	return nil
}
