package repository

import (
	"context"
	"testing"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/pkg/database"
	"github.com/stretchr/testify/require"
)

func TestTeacherRepositoryE2E(t *testing.T) {
	db, err := database.NewPostgres(database.Config{
		DSN:             TestDBConnString,
		MaxOpenConns:    200,
		MaxIdleConns:    5000,
		ConnMaxLifetime: 600,
	})
	require.NoError(t, err)

	userRepo := NewUserRepository(db)
	teacherRepo := NewTeacherRepository(db)

	ctx := context.Background()

	t.Run("Create Teacher and Get Teacher", func(t *testing.T) {
		teacher := &domain.Teacher{
			User: domain.User{
				Name:     "John Doe",
				Email:    "johndoe@gmail.com",
				Password: "Skibidi12345",
			},
		}

		err := teacherRepo.Create(ctx, teacher)
		require.NoError(t, err)

		// Check if teacher does exist in db by email
		user, err := userRepo.GetByEmail(ctx, teacher.Email)
		require.NoError(t, err)
		require.Equal(t, user.Email, teacher.Email)

		// Check if teacher does exist in db by ID
		_, err = userRepo.GetByID(ctx, teacher.ID)
		require.NoError(t, err)
	})
}
