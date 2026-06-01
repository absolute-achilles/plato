package repository

import (
	"context"
	"fmt"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/dto"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	Create(ctx context.Context, req *dto.CreateUserRequest) (*domain.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	var user domain.User
	err := r.db.GetContext(ctx, &user, `SELECT * FROM users WHERE id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByID: %w", err)
	}
	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, req *dto.CreateUserRequest) (*domain.User, error) {
	var user domain.User
	query := `
        INSERT INTO users (name, email, role)
        VALUES ($1, $2, $3)
        RETURNING id, name, email, role, created_at`
	err := r.db.QueryRowxContext(ctx, query, req.Name, req.Email, req.Role).StructScan(&user)
	if err != nil {
		return nil, fmt.Errorf("userRepository.Create: %w", err)
	}
	return &user, nil
}
