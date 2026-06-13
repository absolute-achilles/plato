package repository

import (
	"context"
	"testing"
	"time"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/pkg/common"
	"github.com/absolute-achilles/plato/pkg/database"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

func TestStudentRepositoryE2E(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	postgresC, err := createPostgresContainer(ctx)
	require.NoError(t, err, "failed to setup postgres test container")
	defer testcontainers.CleanupContainer(t, postgresC)

	connStr, err := postgresC.ConnectionString(ctx, "sslmode=disable")

	require.NoError(t, err, "could not get DB connection string")
	db, err := database.NewPostgres(database.Config{
		DSN:             connStr,
		MaxConns:        common.Int32Ptr(200),
		MinIdleConns:    common.Int32Ptr(5000),
		ConnMaxLifetime: common.TimeDurationPtr(10 * time.Minute),
	})
	require.NoError(t, err)

	userRepo := NewUserRepository(db)

	studentRepo := NewStudentRepository(db)

	t.Run("Create Student and Get Student", func(t *testing.T) {
		t.Parallel()
		student := &domain.Student{
			User: domain.User{
				Username:     "John Doe Student",
				Email:        "student@gmail.com",
				HashPassword: "Skibidi12345",
			},
		}

		err := studentRepo.Create(ctx, student)
		require.NoError(t, err)

		// Check if student does exist in db by email
		user, err := userRepo.GetByEmail(ctx, student.Email)
		require.NoError(t, err)
		require.Equal(t, user.Email, student.Email)
		require.Equal(t, user.Username, student.Username)

		// Check if student does exist in db by ID
		user, err = userRepo.GetByID(ctx, student.ID)
		require.NoError(t, err)
		require.Equal(t, user.Email, student.Email)
		require.Equal(t, user.Username, student.Username)
	})

}
