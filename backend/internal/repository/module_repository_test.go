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

func TestModuleRepository(t *testing.T) {
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

	teacherRepo := NewTeacherRepository(db)
	courseRepo := NewCourseRepository(db)
	moduleRepo := NewModuleRepository(db)

	teacher := &domain.Teacher{
		User: domain.User{
			Username:     "John Snow",
			Email:        "johnsnow@gmail.com",
			HashPassword: "Skibidi12345",
		},
	}

	err = teacherRepo.Create(ctx, teacher)
	require.NoError(t, err)

	t.Run("Create module in a course and get it", func(t *testing.T) {
		t.Parallel()

		course := &domain.Course{
			TeacherID:   teacher.ID,
			Name:        "Game of Thrones",
			Description: "About game of thrones",
		}

		courseID, err := courseRepo.Create(ctx, course)
		require.NoError(t, err)

		_, err = courseRepo.Get(ctx, courseID)
		require.NoError(t, err, "failed to get course")

		module := &domain.Module{
			CourseID: courseID,
			Name:     "Week 1",
			Position: 400,
		}
		moduleID, err := moduleRepo.Create(ctx, module)
		require.NoError(t, err, "failed to create module")

		_, err = moduleRepo.Get(ctx, moduleID)
		require.NoError(t, err, "failed to get module")
	})

	t.Run("Get Course modules", func(t *testing.T) {
		t.Parallel()

		course := &domain.Course{
			TeacherID:   teacher.ID,
			Name:        "Return of the king",
			Description: "About game of thrones",
		}

		courseID, err := courseRepo.Create(ctx, course)
		require.NoError(t, err)

		expectedModules := []*domain.Module{
			{
				CourseID: courseID,
				Name:     "The Great Battle",
				Position: 400,
			},
			{
				CourseID: courseID,
				Name:     "Aragon The King",
				Position: 800,
			},
		}

		_, err = courseRepo.Get(ctx, courseID)
		require.NoError(t, err, "failed to get course")

		for i, module := range expectedModules {
			moduleID, err := moduleRepo.Create(ctx, module)
			expectedModules[i].ID = moduleID
			require.NoError(t, err, "failed to create module")
		}

		dataModules, err := moduleRepo.GetCourseModules(ctx, courseID)
		require.NoError(t, err, "failed to get course modules")

		for _, expected := range expectedModules {
			found := false
			for _, actual := range dataModules {
				if actual.ID == expected.ID {
					require.Equal(t, expected.Name, actual.Name)
					require.Equal(t, expected.Position, actual.Position)
					found = true
					break
				}
			}
			require.True(t, found, "expected module not found in database results")
		}
	})
}
