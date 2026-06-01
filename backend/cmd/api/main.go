package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/absolute-achilles/plato/internal/middleware"
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
	// Handle if we should fail the program if the db connection fails. Or program should continue if db connection failed

	// db, err := database.NewPostgres(database.Config{
	// 	DSN:             os.Getenv("DATABASE_URL"),
	// 	MaxOpenConns:    25,
	// 	MaxIdleConns:    10,
	// 	ConnMaxLifetime: 5 * time.Minute,
	// })
	// if err != nil {
	// 	slog.Error("failed to connect to database", "error", err)
	// 	os.Exit(1)
	// }
	// defer db.Close()

	// 2. Create repositories
	// 3. Create services
	// 4. Create handlers

	// userHandler := NewUserHandler(userSvc)

	// 5. create gin
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.TimeoutMiddleware(Timeout))
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// 7. assign gin handlers
	_ = r.Group("/api/v1")
	// api := r.Group("/api/v1")

	slog.Info("server starting", "port", "8080")
	if err := r.Run(":8080"); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
	// userHandler.RegisterRoutes(api)
}
