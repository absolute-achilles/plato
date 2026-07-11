package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/absolute-achilles/plato/internal/handler"
	"github.com/absolute-achilles/plato/internal/middleware"
	"github.com/absolute-achilles/plato/internal/repository"
	"github.com/absolute-achilles/plato/internal/service"
	"github.com/absolute-achilles/plato/pkg/database"
	"github.com/absolute-achilles/plato/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	// hmmmm
	Timeout time.Duration = 10 * time.Second
)

func main() {
	// Setup logger
	log, cleanup, err := logger.New(logger.Config{
		Level:     slog.LevelInfo,
		LogDir:    "logs",
		Filename:  "app.log",
		AddSource: true,
	})
	if err != nil {
		// can't use our logger yet, fall back to stdlib
		slog.Error("failed to init logger", "error", err)
		os.Exit(1)
	}
	defer cleanup()

	slog.SetDefault(log)

	log.Info("Starting application")

	// 1. Create DB connection
	maxConns := int32(25)
	minIdleConns := int32(10)
	connMaxLifetime := 5 * time.Minute

	db, err := database.NewPostgres(database.Config{
		DSN:             os.Getenv("DATABASE_URL"),
		MaxConns:        &maxConns,
		MinIdleConns:    &minIdleConns,
		ConnMaxLifetime: &connMaxLifetime,
	})
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// 2. Create repositories
	userRepo := repository.NewUserRepository(db)
	teacherRepo := repository.NewTeacherRepository(db)
	studentRepo := repository.NewStudentRepository(db)
	parentRepo := repository.NewParentRepository(db)

	// 3. Create services
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret-change-in-production"
		slog.Warn("JWT_SECRET not set; using insecure default secret")
	}
	authCfg := service.DefaultAuthConfig(jwtSecret)
	authSvc := service.NewAuthService(userRepo, authCfg)
	adminSvc := service.NewAdminService(teacherRepo, studentRepo, parentRepo)

	// 4. Create handlers
	authHandler := handler.NewAuthHandler(authSvc)
	adminHandler := handler.NewAdminHandler(adminSvc, authSvc)

	// 5. create gin
	r := gin.New()
	r.Use(gin.Recovery())
	r.SetTrustedProxies(nil)
	r.Use(middleware.TimeoutMiddleware(Timeout))
	r.Use(middleware.RateLimiterMiddleware())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// 7. assign gin handlers
	api := r.Group("/api/v1")
	handler.RegisterHealthCheck(api)
	authHandler.RegisterRoutes(api)
	adminHandler.RegisterRoutes(api)

	slog.Info("server starting", "port", "8080")
	if err := r.Run(":8080"); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
