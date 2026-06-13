package service

import (
	"testing"
)

func TestContentAttachmentService(t *testing.T) {
	// t.Parallel()
	// ctx := context.Background()
	//
	// postgresC, err := createPostgresContainer(ctx)
	// require.NoError(t, err, "failed to setup postgres test container")
	// defer testcontainers.CleanupContainer(t, postgresC)
	//
	// connStr, err := postgresC.ConnectionString(ctx, "sslmode=disable")
	//
	// require.NoError(t, err, "could not get DB connection string")
	// db, err := database.NewPostgres(database.Config{
	// 	DSN:             connStr,
	// 	MaxConns:        common.Int32Ptr(200),
	// 	MinIdleConns:    common.Int32Ptr(5000),
	// 	ConnMaxLifetime: common.TimeDurationPtr(10 * time.Minute),
	// })
	//
	// teacherRepo := repository.NewTeacherRepository(db)
	// courseRepo := repository.NewCourseRepository(db)
	// moduleRepo := repository.NewModuleRepository(db)
	// moduleContentRepo := repository.NewModuleContentRepository(db)
	// contentAttachmentRepo := repository.NewContentAttachmentRepository(db)
	//
	// st := storage.NewMockStorage()
	// contentAttachmentService := NewContentAttachmentService(contentAttachmentRepo)
}
