package service

import (
	"testing"

	"github.com/absolute-achilles/plato/internal/repository"
	"github.com/absolute-achilles/plato/pkg/database"
	"github.com/stretchr/testify/require"
)

func TestUserServiceE2E(t *testing.T) {
	connStr, cleanp := createPosgresContainer(t)
	defer cleanp()

	db, err := database.NewPostgres(database.Config{
		DSN:             connStr,
		MaxOpenConns:    200,
		MaxIdleConns:    5000,
		ConnMaxLifetime: 600,
	})
	require.NoError(t, err)

	userRepo := repository.NewUserRepository(db)
	_ = NewUserService(userRepo)

	// userService.CreateUser(ctx)

}
