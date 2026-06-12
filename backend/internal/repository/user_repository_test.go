package repository

import (
	"context"
	"testing"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/pkg/database"
	"github.com/stretchr/testify/require"
)

func TestUserRepositoryE2E(t *testing.T) {
	db, err := database.NewPostgres(database.Config{
		DSN:             TestDBConnString,
		MaxOpenConns:    200,
		MaxIdleConns:    5000,
		ConnMaxLifetime: 600,
	})
	require.NoError(t, err)

	userRepo := NewUserRepository(db)

	// Test using student repository
	studentRepo := NewStudentRepository(db)

	ctx := context.Background()

	t.Run("Create User and Get User", func(t *testing.T) {
		t.Parallel()
		student := &domain.Student{
			User: domain.User{
				Name:         "Jesse Pinkman",
				Email:        "jessepink@gmail.com",
				HashPassword: "Skibidi12345",
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

	t.Run("Check duplicate email user", func(t *testing.T) {
		student1 := &domain.Student{
			User: domain.User{
				Name:         "Walter White",
				Email:        "walterwhite@gmail.com",
				HashPassword: "SkibidiToilet",
			},
		}

		student2 := &domain.Student{
			User: domain.User{
				Name:         "Walter Green",
				Email:        "walterwhite@gmail.com",
				HashPassword: "SkibidiToilet",
			},
		}

		ctx := context.Background()
		err := studentRepo.Create(ctx, student1)
		require.NoError(t, err)

		err = studentRepo.Create(ctx, student2)
		require.Error(t, err)
	})

	t.Run("Check duplicate username user", func(t *testing.T) {
		t.Parallel()
		student1 := &domain.Student{
			User: domain.User{
				Name:         "Johnny Blaze",
				Email:        "johnnyblaze456@gmail.com",
				HashPassword: "SkibidiToilet",
			},
		}

		student2 := &domain.Student{
			User: domain.User{
				Name:         "Johnny Blaze",
				Email:        "johnnyblaze123@gmail.com",
				HashPassword: "SkibidiToilet",
			},
		}

		ctx := context.Background()
		err := studentRepo.Create(ctx, student1)
		require.NoError(t, err)

		err = studentRepo.Create(ctx, student2)
		require.Error(t, err)
	})

	t.Run("Change user name", func(t *testing.T) {
		t.Parallel()

		student := &domain.Student{
			User: domain.User{
				Name:         "Spiderman",
				Email:        "spiderman@gmail.com",
				HashPassword: "SkibidiToilet",
			},
		}

		ctx := context.Background()
		err := studentRepo.Create(ctx, student)
		require.NoError(t, err)

		err = userRepo.ChangeName(ctx, student.ID, "Senior Spiderman")
		require.NoError(t, err)
	})

	t.Run("Change user name (duplicate)", func(t *testing.T) {
		t.Parallel()

		student1 := &domain.Student{
			User: domain.User{
				Name:         "Iron Man",
				Email:        "ironman@gmail.com",
				HashPassword: "SkibidiToilet",
			},
		}

		student2 := &domain.Student{
			User: domain.User{
				Name:         "Hulk",
				Email:        "hulk@gmail.com",
				HashPassword: "SkibidiToilet",
			},
		}

		ctx := context.Background()
		err := studentRepo.Create(ctx, student1)
		require.NoError(t, err)

		err = studentRepo.Create(ctx, student2)
		require.NoError(t, err)

		err = userRepo.ChangeName(ctx, student1.ID, "Hulk")
		require.Error(t, err)
	})

	t.Run("Change password", func(t *testing.T) {
		t.Parallel()

		student := &domain.Student{
			User: domain.User{
				Name:         "Captain America",
				Email:        "captainamerica@gmail.com",
				HashPassword: "SkibidiToilet",
			},
		}

		ctx := context.Background()
		err := studentRepo.Create(ctx, student)
		require.NoError(t, err)

		// Wrong old password
		err = userRepo.ChangePassword(ctx, student.ID, "Wrong Password", "Hulk")
		require.ErrorAs(t, err, &ErrPasswordDoesNotMatch)

		// correct old password
		err = userRepo.ChangePassword(ctx, student.ID, "SkibidiToilet", "Hulk")
		require.NoError(t, err)
	})

}
