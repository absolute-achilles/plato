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

func TestCourseRepository(t *testing.T) {
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

	t.Run("Create course and Get course", func(t *testing.T) {
		t.Parallel()
		teacher := &domain.Teacher{
			User: domain.User{
				Username:     "John Snow",
				Email:        "johnsnow@gmail.com",
				HashPassword: "Skibidi12345",
			},
		}

		err := teacherRepo.Create(ctx, teacher)
		require.NoError(t, err)

		course := &domain.Course{
			TeacherID:   teacher.ID,
			Name:        "Game of Thrones",
			Description: "About game of thrones",
		}

		err = courseRepo.Create(ctx, course)
		require.NoError(t, err)

		// Duplicate course is allowed here (in the meantime)
		err = courseRepo.Create(ctx, course)
		require.NoError(t, err)
	})

	t.Run("Teacher Courses", func(t *testing.T) {
		t.Parallel()
		teacher := &domain.Teacher{
			User: domain.User{
				Username:     "Tywin Lannister",
				Email:        "tywinlannister@gmail.com",
				HashPassword: "Skibidi12345",
			},
		}

		err := teacherRepo.Create(ctx, teacher)
		require.NoError(t, err)

		teacherCourses := []*domain.Course{
			{
				TeacherID:   teacher.ID,
				Name:        "History of Lannister",
				Description: "About",
			},
			{
				TeacherID:   teacher.ID,
				Name:        "History of Winterfell",
				Description: "About",
			},
		}

		for _, course := range teacherCourses {
			err = courseRepo.Create(ctx, course)
			require.NoError(t, err)
		}

		dataCourses, err := courseRepo.GetTeacherCoursesByID(ctx, teacher.ID)
		require.NoError(t, err)
		require.Len(t, dataCourses, len(teacherCourses))

		for _, expected := range teacherCourses {
			found := false
			for _, actual := range dataCourses {
				if actual.ID == expected.ID {
					require.Equal(t, expected.Name, actual.Name)
					require.Equal(t, expected.Description, actual.Description)
					found = true
					break
				}
			}
			require.True(t, found, "expected course not found in database results")
		}

		dataCourses, err = courseRepo.GetTeacherCoursesByEmail(ctx, teacher.Email)
		require.NoError(t, err)
		require.Len(t, dataCourses, len(teacherCourses))

		for _, expected := range teacherCourses {
			found := false
			for _, actual := range dataCourses {
				if actual.ID == expected.ID {
					require.Equal(t, expected.Name, actual.Name)
					require.Equal(t, expected.Description, actual.Description)
					found = true
					break
				}
			}
			require.True(t, found, "expected course not found in database results")
		}

	})
}
