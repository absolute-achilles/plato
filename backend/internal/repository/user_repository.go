package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	query := `SELECT id, username, email, hash_password, role, created_at FROM users WHERE id = $1`

	row, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByID (query): %w", err)
	}

	user, err := pgx.CollectExactlyOneRow(row, pgx.RowToAddrOfStructByName[domain.User])
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByID (collect): %w", err)
	}

	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, username, email, hash_password, role, created_at FROM users WHERE email = $1`
	rows, err := r.db.Query(ctx, query, email)
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByEmail: %w", err)
	}
	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.User])
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByEmail: %w", err)
	}
	return user, nil
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

	result, err := r.db.Exec(ctx, queryUpdatePassword, newPassword, id)
	if err != nil {
		return fmt.Errorf("userRepository.ChangePassword: %w", err)
	}

	if result.RowsAffected() == 0 {
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
	result, err := r.db.Exec(ctx, query, newName, id)
	if err != nil {
		return fmt.Errorf("userRepository.ChangeName: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user not found with ID: %s", id)
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("userRepository.Delete: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("userRepository.Delete: user not found with ID: %s", id)
	}

	return nil
}
