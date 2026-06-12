package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/utils"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	ChangeName(ctx context.Context, id string, newName string) error
	ChangePassword(ctx context.Context, id, oldPassword, newPassword string) error
	Delete(ctx context.Context, id string) error

	// Note: We do NOT have a Create() here. Creation is handled
	// by the specific role repositories to ensure transaction safety.
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, username, email, hash_password, role FROM users WHERE id = $1`
	if err := r.db.GetContext(ctx, &user, query, id); err != nil {
		return nil, fmt.Errorf("userRepository.GetByEmail: %w", err)
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, username, email, hash_password, role FROM users WHERE email = $1`
	if err := r.db.GetContext(ctx, &user, query, email); err != nil {
		return nil, fmt.Errorf("userRepository.GetByEmail: %w", err)
	}
	return &user, nil
}

func (r *userRepository) ChangePassword(ctx context.Context, id, oldPassword, newPassword string) error {
	user, err := r.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if !utils.CheckHashPassword(oldPassword, user.HashPassword) {
		return ErrPasswordDoesNotMatch
	}

	queryUpdatePassword := `
	UPDATE users
	SET hash_password = $1
	WHERE id = $2
	`

	result, err := r.db.ExecContext(ctx, queryUpdatePassword, newPassword, id)
	if err != nil {
		return fmt.Errorf("userRepository.ChangePassword: %w", err)
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

func (r *userRepository) ChangeName(ctx context.Context, id string, newName string) error {
	query := `
	UPDATE users
	SET username = $1
	WHERE id = $2
	`
	result, err := r.db.ExecContext(ctx, query, newName, id)
	if err != nil {
		return fmt.Errorf("userRepository.ChangeName: %w", err)
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

func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`

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
