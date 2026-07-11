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
		MaxConns:        common.Int32Ptr(200),
		MinIdleConns:    common.Int32Ptr(5000),
		ConnMaxLifetime: common.TimeDurationPtr(10 * time.Minute),
	})
	require.NoError(t, err)

	userRepo := NewUserRepository(db)
	studentRepo := NewStudentRepository(db)
	parentRepo := NewParentRepository(db)

	t.Run("Create Parent and Get Parent", func(t *testing.T) {
		t.Parallel()

		student := &domain.Student{
			User: domain.User{
				Username:     "Student Child",
				Email:        "studentchild@gmail.com",
				HashPassword: "Skibidi12345",
			},
			GradeLevel: domain.GradeLevel5,
		}
		err := studentRepo.Create(ctx, student)
		require.NoError(t, err)

		parent := &domain.Parent{
			User: domain.User{
				Username:     "John Doe Parent",
				Email:        "parent@gmail.com",
				HashPassword: "Skibidi12345",
			},
			Type:       domain.ParentRelationshipFather,
			StudentIDs: []string{student.ID},
		}

		err = parentRepo.Create(ctx, parent)
		require.NoError(t, err)

		// Check if parent exists in db by email
		user, err := userRepo.GetByEmail(ctx, parent.Email)
		require.NoError(t, err)
		require.Equal(t, user.Email, parent.Email)
		require.Equal(t, user.Username, parent.Username)

		// Check if parent exists in db by ID
		user, err = userRepo.GetByID(ctx, parent.ID)
		require.NoError(t, err)
		require.Equal(t, user.Email, parent.Email)
		require.Equal(t, user.Username, parent.Username)
	})

}
