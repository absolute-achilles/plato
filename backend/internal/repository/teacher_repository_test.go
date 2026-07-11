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

func TestTeacherRepositoryE2E(t *testing.T) {
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

	teacherRepo := NewTeacherRepository(db)

	t.Run("Create Teacher and Get Teacher", func(t *testing.T) {
		t.Parallel()
	teacher := &domain.Teacher{
		User: domain.User{
			Username:     "John Doe Teacher",
			Email:        "teacher@gmail.com",
			HashPassword: "Skibidi12345",
		},
		Department: domain.DepartmentMathematics,
	}

		err := teacherRepo.Create(ctx, teacher)
		require.NoError(t, err)

		// Check if student does exist in db by email
		user, err := userRepo.GetByEmail(ctx, teacher.Email)
		require.NoError(t, err)
		require.Equal(t, user.Email, teacher.Email)
		require.Equal(t, user.Username, teacher.Username)

		// Check if student does exist in db by ID
		user, err = userRepo.GetByID(ctx, teacher.ID)
		require.NoError(t, err)
		require.Equal(t, user.Email, teacher.Email)
		require.Equal(t, user.Username, teacher.Username)
	})

}
