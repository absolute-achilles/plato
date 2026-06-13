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

func TestEnrollmentRepository(t *testing.T) {
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

	teacherRepo := NewTeacherRepository(db)
	studentRepo := NewStudentRepository(db)
	courseRepo := NewCourseRepository(db)
	enrollmentRepo := NewEnrollmentRepository(db)

	// Create a shared teacher for all courses in this test suite
	teacher := &domain.Teacher{
		User: domain.User{
			Username:     "Ned Stark",
			Email:        "ned@winterfell.com",
			HashPassword: "password123",
		},
	}
	err = teacherRepo.Create(ctx, teacher)
	require.NoError(t, err)

	t.Run("Create and Get Enrollment", func(t *testing.T) {
		t.Parallel()

		// 1. Setup isolated data
		student := &domain.Student{
			User: domain.User{
				Username:     "Arya Stark",
				Email:        "arya@stark.com",
				HashPassword: "123",
			},
		}
		err := studentRepo.Create(ctx, student)
		require.NoError(t, err)

		course := &domain.Course{
			TeacherID:   teacher.ID,
			Name:        "Water Dancing 101",
			Description: "Sword fighting basics",
		}
		courseID, err := courseRepo.Create(ctx, course)
		require.NoError(t, err)

		expected := &domain.Enrollment{
			StudentID: student.ID,
			CourseID:  courseID,
		}
		enrollmentID, err := enrollmentRepo.Create(ctx, expected)
		require.NoError(t, err)

		actual, err := enrollmentRepo.Get(ctx, enrollmentID)
		require.NoError(t, err)
		require.Equal(t, enrollmentID, actual.ID)
		require.Equal(t, expected.StudentID, actual.StudentID)
		require.Equal(t, expected.CourseID, actual.CourseID)
	})

	t.Run("Get By Student And Course", func(t *testing.T) {
		t.Parallel()

		student := &domain.Student{
			User: domain.User{Username: "Robb Stark", Email: "robb@stark.com", HashPassword: "123"},
		}
		err := studentRepo.Create(ctx, student)
		require.NoError(t, err)

		course := &domain.Course{
			TeacherID:   teacher.ID,
			Name:        "Battle Tactics",
			Description: "Winning the war",
		}
		courseID, err := courseRepo.Create(ctx, course)
		require.NoError(t, err)

		enrollmentID, err := enrollmentRepo.Create(ctx, &domain.Enrollment{
			StudentID: student.ID,
			CourseID:  courseID,
		})
		require.NoError(t, err)

		actual, err := enrollmentRepo.GetByStudentAndCourse(ctx, student.ID, courseID)
		require.NoError(t, err)
		require.Equal(t, enrollmentID, actual.ID)
	})

	t.Run("Get Student Enrollments", func(t *testing.T) {
		t.Parallel()

		student := &domain.Student{
			User: domain.User{
				Username:     "Jon Snow",
				Email:        "jon@wall.com",
				HashPassword: "123",
			},
		}
		err := studentRepo.Create(ctx, student)
		require.NoError(t, err)

		course1 := &domain.Course{TeacherID: teacher.ID, Name: "Wall Defense 1"}
		course1ID, _ := courseRepo.Create(ctx, course1)

		course2 := &domain.Course{TeacherID: teacher.ID, Name: "Wildling Relations"}
		course2ID, _ := courseRepo.Create(ctx, course2)

		_, err = enrollmentRepo.Create(ctx, &domain.Enrollment{StudentID: student.ID, CourseID: course1ID})
		require.NoError(t, err)

		_, err = enrollmentRepo.Create(ctx, &domain.Enrollment{StudentID: student.ID, CourseID: course2ID})
		require.NoError(t, err)

		enrollments, err := enrollmentRepo.GetStudentEnrollments(ctx, student.ID)
		require.NoError(t, err)
		require.Len(t, enrollments, 2)
	})

	t.Run("Get Course Enrollments", func(t *testing.T) {
		t.Parallel()

		course := &domain.Course{
			TeacherID: teacher.ID,
			Name:      "History of Westeros",
		}
		courseID, err := courseRepo.Create(ctx, course)
		require.NoError(t, err)

		student1 := &domain.Student{
			User: domain.User{Username: "Sansa Stark", Email: "sansa@stark.com", HashPassword: "123"},
		}
		studentRepo.Create(ctx, student1)

		student2 := &domain.Student{
			User: domain.User{Username: "Bran Stark", Email: "bran@stark.com", HashPassword: "123"},
		}
		studentRepo.Create(ctx, student2)

		_, err = enrollmentRepo.Create(ctx, &domain.Enrollment{StudentID: student1.ID, CourseID: courseID})
		require.NoError(t, err)

		_, err = enrollmentRepo.Create(ctx, &domain.Enrollment{StudentID: student2.ID, CourseID: courseID})
		require.NoError(t, err)

		enrollments, err := enrollmentRepo.GetCourseEnrollments(ctx, courseID)
		require.NoError(t, err)
		require.Len(t, enrollments, 2)
	})

	t.Run("Delete Enrollment", func(t *testing.T) {
		t.Parallel()

		student := &domain.Student{
			User: domain.User{
				Username:     "Theon Greyjoy",
				Email:        "theon@ironislands.com",
				HashPassword: "123",
			},
		}
		err := studentRepo.Create(ctx, student)
		require.NoError(t, err)

		course := &domain.Course{
			TeacherID: teacher.ID,
			Name:      "Archery 101",
		}
		courseID, err := courseRepo.Create(ctx, course)
		require.NoError(t, err)

		enrollmentID, err := enrollmentRepo.Create(ctx, &domain.Enrollment{
			StudentID: student.ID,
			CourseID:  courseID,
		})
		require.NoError(t, err)

		err = enrollmentRepo.Delete(ctx, enrollmentID)
		require.NoError(t, err)

		// Verify it was deleted
		_, err = enrollmentRepo.Get(ctx, enrollmentID)
		require.Error(t, err)
		require.Contains(t, err.Error(), "no rows in result set")
	})
}
