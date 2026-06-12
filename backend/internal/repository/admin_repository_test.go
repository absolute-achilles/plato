package repository

import (
	"context"
	"testing"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/pkg/database"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

func TestAdminRepositoryE2E(t *testing.T) {
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

	adminRepo := NewAdminRepository(db)

	t.Run("Create Admin and Get Admin", func(t *testing.T) {
		t.Parallel()
		admin := &domain.Admin{
			User: domain.User{
				Username:     "John Doe Teacher",
				Email:        "student@gmail.com",
				HashPassword: "Skibidi12345",
			},
		}

		err := adminRepo.Create(ctx, admin)
		require.NoError(t, err)

		// Check if student does exist in db by email
		user, err := userRepo.GetByEmail(ctx, admin.Email)
		require.NoError(t, err)
		require.Equal(t, user.Email, admin.Email)
		require.Equal(t, user.Username, admin.Username)

		// Check if student does exist in db by ID
		user, err = userRepo.GetByID(ctx, admin.ID)
		require.NoError(t, err)
		require.Equal(t, user.Email, admin.Email)
		require.Equal(t, user.Username, admin.Username)
	})

}
