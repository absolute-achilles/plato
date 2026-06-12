package repository

import (
	"context"
	"testing"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/pkg/database"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

func TestParentRepositoryE2E(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	postgresC, err := createPostgresContainer(ctx)
	require.NoError(t, err, "failed to setup postgres test container")
	defer testcontainers.CleanupContainer(t, postgresC)

	connStr, err := postgresC.ConnectionString(ctx, "sslmode=disable")

	require.NoError(t, err, "could not get DB connection string")
	db, err := database.NewPostgres(database.Config{
		DSN:             connStr,
		MaxOpenConns:    200,
		MaxIdleConns:    5000,
		ConnMaxLifetime: 600,
	})
	require.NoError(t, err)

	userRepo := NewUserRepository(db)

	parentRepo := NewParentRepository(db)

	t.Run("Create Parent and Get Parent", func(t *testing.T) {
		t.Parallel()
		parent := &domain.Parent{
			User: domain.User{
				Name:         "John Doe Teacher",
				Email:        "student@gmail.com",
				HashPassword: "Skibidi12345",
			},
		}

		err := parentRepo.Create(ctx, parent)
		require.NoError(t, err)

		// Check if student does exist in db by email
		user, err := userRepo.GetByEmail(ctx, parent.Email)
		require.NoError(t, err)
		require.Equal(t, user.Email, parent.Email)

		// Check if student does exist in db by ID
		_, err = userRepo.GetByID(ctx, parent.ID)
		require.NoError(t, err)
	})

}
