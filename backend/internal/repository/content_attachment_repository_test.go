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

func TestContentAttachmentRepository(t *testing.T) {
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
	contentAttachmentRepo := NewContentAttachmentRepository(db)

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

	module := &domain.Module{
		CourseID: courseID,
		Name:     "Week 1",
		Position: 400,
	}
	moduleID, err := moduleRepo.Create(ctx, module)
	require.NoError(t, err, "failed to create module")

	_, err = moduleRepo.Get(ctx, moduleID)
	require.NoError(t, err, "failed to get module")

	moduleContent := &domain.ModuleContent{
		ModuleID: moduleID,
		Title:    "The Starks",
		// Markdown or HTML
		BodyContent:       "## History of the starks\nThe Starks are one of the power families in the game of throne world",
		ModuleContentType: domain.ModuleContentTypeLesson,
		Position:          400,
	}

	moduleContentID, err := moduleContentRepo.Create(ctx, moduleContent)
	require.NoError(t, err)

	_, err = moduleContentRepo.Get(ctx, moduleContentID)
	require.NoError(t, err)

	t.Run("Create content attachment in a module content and get it", func(t *testing.T) {
		t.Parallel()

		expected := &domain.ContentAttachment{
			ModuleContentID: moduleContentID,
			Name:            "Seven Kingdoms Map",
			URL:             "https://somewhere.com/map.png",
			SizeBytes:       1000,
			FileType:        domain.FileTypeImagePng,
		}

		contentAttachmentID, err := contentAttachmentRepo.Create(ctx, expected)
		require.NoError(t, err)

		actual, err := contentAttachmentRepo.Get(ctx, contentAttachmentID)
		require.NoError(t, err)

		require.Equal(t, actual.Name, expected.Name)
		require.Equal(t, actual.URL, expected.URL)
		require.Equal(t, actual.SizeBytes, expected.SizeBytes)
		require.Equal(t, actual.FileType, expected.FileType)
	})

	t.Run("Get Module Content Attachments", func(t *testing.T) {
		t.Parallel()

		// 1. Create a dedicated Module Content just for this subtest
		// to ensure data isolation from other parallel tests.
		testModuleContent := &domain.ModuleContent{
			ModuleID:          moduleID, // Reusing the module created in the main test setup
			Title:             "The White Walkers",
			BodyContent:       "## Threats beyond the wall\nWinter is coming.",
			ModuleContentType: domain.ModuleContentTypeLesson,
			Position:          500,
		}

		testModuleContentID, err := moduleContentRepo.Create(ctx, testModuleContent)
		require.NoError(t, err)

		// 2. Define the attachments we expect to insert and retrieve
		expectedAttachments := []*domain.ContentAttachment{
			{
				ModuleContentID: testModuleContentID,
				Name:            "White Walker Anatomy",
				URL:             "https://somewhere.com/whitewalker.png",
				SizeBytes:       2048,
				FileType:        domain.FileTypeImagePng,
			},
			{
				ModuleContentID: testModuleContentID,
				Name:            "Battle of Hardhome",
				URL:             "https://somewhere.com/hardhome.mp4",
				SizeBytes:       102400,
				FileType:        domain.FileTypeVideoMp4,
			},
		}

		for i, attachment := range expectedAttachments {
			id, err := contentAttachmentRepo.Create(ctx, attachment)
			require.NoError(t, err, "failed to create content attachment")
			expectedAttachments[i].ID = id
		}

		actualAttachments, err := contentAttachmentRepo.GetModuleContentAttachments(ctx, testModuleContentID)
		require.NoError(t, err, "failed to get module content attachments")

		require.Len(t, actualAttachments, len(expectedAttachments), "should return exactly the number of inserted attachments")

		for _, expected := range expectedAttachments {
			found := false
			for _, actual := range actualAttachments {
				if actual.ID == expected.ID {
					require.Equal(t, expected.Name, actual.Name)
					require.Equal(t, expected.URL, actual.URL)
					require.Equal(t, expected.SizeBytes, actual.SizeBytes)
					require.Equal(t, expected.FileType, actual.FileType)
					require.Equal(t, expected.ModuleContentID, actual.ModuleContentID)
					found = true
					break
				}
			}
			require.True(t, found, "expected attachment '%s' not found in database results", expected.Name)
		}
	})
}
