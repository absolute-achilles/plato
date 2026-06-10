package repository

import (
	"context"
	"testing"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/pkg/database"
	"github.com/stretchr/testify/require"
)

func TestStudentRepositoryE2E(t *testing.T) {
	db, err := database.NewPostgres(database.Config{
		DSN:             TestDBConnString,
		MaxOpenConns:    200,
		MaxIdleConns:    5000,
		ConnMaxLifetime: 600,
	})
	require.NoError(t, err)

	userRepo := NewUserRepository(db)
	studentRepo := NewStudentRepository(db)

	ctx := context.Background()

	t.Run("Create Teacher and Get Teacher", func(t *testing.T) {
		student := &domain.Student{
			User: domain.User{
				Name:     "Jesse Pinkman",
				Email:    "jessepink@gmail.com",
				Password: "Skibidi12345",
			},
		}

		err := studentRepo.Create(ctx, student)
		require.NoError(t, err)

		// Check if student does exist in db by email
		user, err := userRepo.GetByEmail(ctx, student.Email)
		require.NoError(t, err)
		require.Equal(t, user.Email, student.Email)

		// Check if student does exist in db by ID
		_, err = userRepo.GetByID(ctx, student.ID)
		require.NoError(t, err)
	})
}
