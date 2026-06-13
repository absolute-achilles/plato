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

func TestAttendanceRepository(t *testing.T) {
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
	moduleRepo := NewModuleRepository(db)
	attendanceRepo := NewAttendanceRepository(db)

	// Global teacher for the test suite
	teacher := &domain.Teacher{
		User: domain.User{
			Username:     "Stannis Baratheon",
			Email:        "stannis@dragonstone.com",
			HashPassword: "password123",
		},
	}
	err = teacherRepo.Create(ctx, teacher)
	require.NoError(t, err)

	t.Run("Create, Get, and Update Attendance", func(t *testing.T) {
		t.Parallel()

		student := &domain.Student{User: domain.User{Username: "Shireen", Email: "shireen@dragonstone.com", HashPassword: "123"}}
		err := studentRepo.Create(ctx, student)
		require.NoError(t, err)

		course := &domain.Course{TeacherID: teacher.ID, Name: "Grammar Rules"}
		courseID, err := courseRepo.Create(ctx, course)
		require.NoError(t, err)

		module := &domain.Module{CourseID: courseID, Name: "Week 1", Position: 1}
		moduleID, err := moduleRepo.Create(ctx, module)
		require.NoError(t, err)

		// Create
		expected := &domain.Attendance{
			StudentID: student.ID,
			ModuleID:  moduleID,
			Status:    domain.AttendanceStatusLate,
			Notes:     common.StringPtr("Arrived 10 minutes late"),
		}

		id, err := attendanceRepo.Create(ctx, expected)
		require.NoError(t, err)

		// Get
		actual, err := attendanceRepo.Get(ctx, id)
		require.NoError(t, err)
		require.Equal(t, id, actual.ID)
		require.Equal(t, expected.Status, actual.Status)
		require.Equal(t, *expected.Notes, *actual.Notes)

		// Update
		newNotes := "Excused by Maester"
		actual.Status = domain.AttendanceStatusExcused
		actual.Notes = &newNotes

		err = attendanceRepo.Update(ctx, actual)
		require.NoError(t, err)

		updated, err := attendanceRepo.Get(ctx, id)
		require.NoError(t, err)
		require.Equal(t, domain.AttendanceStatusExcused, updated.Status)
		require.Equal(t, newNotes, *updated.Notes)
	})

	t.Run("Get Module Attendances", func(t *testing.T) {
		t.Parallel()

		student1 := &domain.Student{User: domain.User{Username: "Davos", Email: "davos@onion.com", HashPassword: "123"}}
		studentRepo.Create(ctx, student1)

		student2 := &domain.Student{User: domain.User{Username: "Melisandre", Email: "melisandre@light.com", HashPassword: "123"}}
		studentRepo.Create(ctx, student2)

		course := &domain.Course{TeacherID: teacher.ID, Name: "Fire Magic"}
		courseID, _ := courseRepo.Create(ctx, course)

		module := &domain.Module{CourseID: courseID, Name: "Introduction", Position: 1}
		moduleID, _ := moduleRepo.Create(ctx, module)

		_, err = attendanceRepo.Create(ctx, &domain.Attendance{StudentID: student1.ID, ModuleID: moduleID, Status: domain.AttendanceStatusPresent})
		require.NoError(t, err)

		_, err = attendanceRepo.Create(ctx, &domain.Attendance{StudentID: student2.ID, ModuleID: moduleID, Status: domain.AttendanceStatusAbsent})
		require.NoError(t, err)

		attendances, err := attendanceRepo.GetModuleAttendances(ctx, moduleID)
		require.NoError(t, err)
		require.Len(t, attendances, 2)
	})

	t.Run("Get By Student And Module", func(t *testing.T) {
		t.Parallel()

		student := &domain.Student{User: domain.User{Username: "Gendry", Email: "gendry@fleabottom.com", HashPassword: "123"}}
		studentRepo.Create(ctx, student)

		course := &domain.Course{TeacherID: teacher.ID, Name: "Smithing"}
		courseID, _ := courseRepo.Create(ctx, course)

		module := &domain.Module{CourseID: courseID, Name: "Hammering", Position: 1}
		moduleID, _ := moduleRepo.Create(ctx, module)

		attendanceID, err := attendanceRepo.Create(ctx, &domain.Attendance{StudentID: student.ID, ModuleID: moduleID, Status: domain.AttendanceStatusPresent})
		require.NoError(t, err)

		actual, err := attendanceRepo.GetByStudentAndModule(ctx, student.ID, moduleID)
		require.NoError(t, err)
		require.Equal(t, attendanceID, actual.ID)
	})

	t.Run("Delete Attendance", func(t *testing.T) {
		t.Parallel()

		student := &domain.Student{User: domain.User{Username: "Salladhor", Email: "sal@pirate.com", HashPassword: "123"}}
		studentRepo.Create(ctx, student)

		course := &domain.Course{TeacherID: teacher.ID, Name: "Naval Tactics"}
		courseID, _ := courseRepo.Create(ctx, course)

		module := &domain.Module{CourseID: courseID, Name: "Boats", Position: 1}
		moduleID, _ := moduleRepo.Create(ctx, module)

		attendanceID, err := attendanceRepo.Create(ctx, &domain.Attendance{
			StudentID: student.ID,
			ModuleID:  moduleID,
			Status:    domain.AttendanceStatusAbsent,
		})
		require.NoError(t, err)

		err = attendanceRepo.Delete(ctx, attendanceID)
		require.NoError(t, err)

		_, err = attendanceRepo.Get(ctx, attendanceID)
		require.Error(t, err)
	})
}
