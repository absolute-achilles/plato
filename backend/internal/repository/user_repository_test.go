package repository

import (
	"context"
	"testing"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/pkg/database"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

func TestUserRepositoryE2E(t *testing.T) {
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

	// Test using student repository
	studentRepo := NewStudentRepository(db)

	t.Run("Create User and Get User", func(t *testing.T) {
		t.Parallel()
		student := &domain.Student{
			User: domain.User{
				Username:     "Jesse Pinkman",
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
		require.Equal(t, user.Username, student.Username)

		// Check if student does exist in db by ID
		user, err = userRepo.GetByID(ctx, student.ID)
		require.NoError(t, err)
		require.Equal(t, user.Username, student.Username)
		require.Equal(t, user.Username, student.Username)
	})

	t.Run("Check duplicate email user", func(t *testing.T) {
		student1 := &domain.Student{
			User: domain.User{
				Username:     "Walter White",
				Email:        "walterwhite@gmail.com",
				HashPassword: "SkibidiToilet",
			},
		}

		student2 := &domain.Student{
			User: domain.User{
				Username:     "Walter Green",
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
				Username:     "Johnny Blaze",
				Email:        "johnnyblaze456@gmail.com",
				HashPassword: "SkibidiToilet",
			},
		}

		student2 := &domain.Student{
			User: domain.User{
				Username:     "Johnny Blaze",
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
				Username:     "Spiderman",
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
				Username:     "Iron Man",
				Email:        "ironman@gmail.com",
				HashPassword: "SkibidiToilet",
			},
		}

		student2 := &domain.Student{
			User: domain.User{
				Username:     "Hulk",
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
				Username:     "Captain America",
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
