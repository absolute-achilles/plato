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

func TestModuleContentRepository(t *testing.T) {
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
	moduleContentRepo := NewModuleContentRepository(db)

	teacher := &domain.Teacher{
		User: domain.User{
			Username:     "John Snow",
			Email:        "johnsnow@gmail.com",
			HashPassword: "Skibidi12345",
		},
	}

	err = teacherRepo.Create(ctx, teacher)
	require.NoError(t, err)

	course := &domain.Course{
		TeacherID:   teacher.ID,
		Name:        "Game of Thrones",
		Description: "About game of thrones",
	}

	courseID, err := courseRepo.Create(ctx, course)
	require.NoError(t, err)

	_, err = courseRepo.Get(ctx, courseID)
	require.NoError(t, err, "failed to get course")

	t.Run("Create module content in a module and get it", func(t *testing.T) {
		t.Parallel()

		module := &domain.Module{
			CourseID: courseID,
			Name:     "Week 1",
			Position: 400,
		}
		moduleID, err := moduleRepo.Create(ctx, module)
		require.NoError(t, err, "failed to create module")

		_, err = moduleRepo.Get(ctx, moduleID)
		require.NoError(t, err, "failed to get module")

		expectedModuleContent := &domain.ModuleContent{
			ModuleID: moduleID,
			Title:    "The Starks",
			// Markdown or HTML
			BodyContent:       "## History of the starks\nThe Starks are one of the power families in the game of throne world",
			ModuleContentType: domain.ModuleContentTypeLesson,
			Position:          400,
		}

		moduleContentID, err := moduleContentRepo.Create(ctx, expectedModuleContent)
		require.NoError(t, err)

		moduleContent, err := moduleContentRepo.Get(ctx, moduleContentID)
		require.NoError(t, err)

		require.Equal(t, moduleContent.Title, expectedModuleContent.Title)
		require.Equal(t, moduleContent.BodyContent, expectedModuleContent.BodyContent)
		require.Equal(t, moduleContent.ModuleID, expectedModuleContent.ModuleID)
	})

	t.Run("Get Module contents", func(t *testing.T) {
		t.Parallel()

		module := &domain.Module{
			CourseID: courseID,
			Name:     "Week 2",
			Position: 600,
		}
		moduleID, err := moduleRepo.Create(ctx, module)
		require.NoError(t, err, "failed to create module")

		_, err = moduleRepo.Get(ctx, moduleID)
		require.NoError(t, err, "failed to get module")

		expectedModuleContents := []*domain.ModuleContent{
			{
				ModuleID: moduleID,
				Title:    "The Lanisters",
				// Markdown or HTML
				BodyContent:       "## History of the Lanisters\nThe Lanisters are one of the power families in the game of throne world",
				ModuleContentType: domain.ModuleContentTypeLesson,
				Position:          500,
			},
			{
				ModuleID: moduleID,
				Title:    "The Targaryen",
				// Markdown or HTML
				BodyContent:       "## History of the Targaryen\nThe Targaryen are one of the power families in the game of throne world",
				ModuleContentType: domain.ModuleContentTypeLesson,
				Position:          600,
			},
		}

		for i, moduleContent := range expectedModuleContents {
			moduleContentID, err := moduleContentRepo.Create(ctx, moduleContent)
			expectedModuleContents[i].ID = moduleContentID
			require.NoError(t, err, "failed to create module content")
		}

		dataModuleContents, err := moduleContentRepo.GetModuleContents(ctx, moduleID)
		require.NoError(t, err, "failed to get module contents")

		for _, expected := range expectedModuleContents {
			found := false
			for _, actual := range dataModuleContents {
				if actual.ID == expected.ID {
					require.Equal(t, expected.Title, actual.Title)
					require.Equal(t, expected.BodyContent, actual.BodyContent)
					require.Equal(t, expected.ModuleContentType, actual.ModuleContentType)
					require.Equal(t, expected.Position, actual.Position)
					found = true
					break
				}
			}
			require.True(t, found, "expected module not found in database results")
		}
	})
}
