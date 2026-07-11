package service

import (
	"context"
	"testing"
	"time"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/dto"
	"github.com/absolute-achilles/plato/internal/repository"
	"github.com/absolute-achilles/plato/internal/utils"
	"github.com/stretchr/testify/require"
)

type mockUserRepository struct {
	users map[string]*domain.User
}

func newMockUserRepository() *mockUserRepository {
	return &mockUserRepository{users: make(map[string]*domain.User)}
}

func (m *mockUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	if u, ok := m.users[id]; ok {
		return u, nil
	}
	return nil, domain.ErrNotFound
}

func (m *mockUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	for _, u := range m.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, domain.ErrNotFound
}

func (m *mockUserRepository) ChangeName(ctx context.Context, id string, newName string) error {
	return nil
}

func (m *mockUserRepository) ChangePassword(ctx context.Context, id, oldPassword, newPassword string) error {
	u, ok := m.users[id]
	if !ok {
		return domain.ErrNotFound
	}
	if !utils.CheckHashPassword(oldPassword, u.HashPassword) {
		return repository.ErrPasswordDoesNotMatch
	}
	hash, _ := utils.HashPassword(newPassword)
	u.HashPassword = hash
	return nil
}

func (m *mockUserRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (m *mockUserRepository) seed(email, password string, role domain.Role) *domain.User {
	hash, _ := utils.HashPassword(password)
	u := &domain.User{
		ID:           "user-1",
		Username:     "admin",
		Email:        email,
		HashPassword: hash,
		Role:         role,
		CreatedAt:    time.Now(),
	}
	m.users[u.ID] = u
	return u
}

func TestAuthService_Login(t *testing.T) {
	ctx := context.Background()
	repo := newMockUserRepository()
	repo.seed("admin@plato.local", "admin12345", domain.RoleAdmin)

	cfg := DefaultAuthConfig("test-secret")
	svc := NewAuthService(repo, cfg)

	t.Run("successful login", func(t *testing.T) {
		resp, err := svc.Login(ctx, &dto.LoginRequest{
			Email:    "admin@plato.local",
			Password: "admin12345",
		})
		require.NoError(t, err)
		require.NotEmpty(t, resp.AccessToken)
		require.NotEmpty(t, resp.RefreshToken)
		require.Equal(t, "admin@plato.local", resp.User.Email)
		require.Equal(t, domain.RoleAdmin, resp.User.Role)
	})

	t.Run("invalid password", func(t *testing.T) {
		_, err := svc.Login(ctx, &dto.LoginRequest{
			Email:    "admin@plato.local",
			Password: "wrong",
		})
		require.ErrorIs(t, err, domain.ErrInvalidInput)
	})

	t.Run("unknown email", func(t *testing.T) {
		_, err := svc.Login(ctx, &dto.LoginRequest{
			Email:    "missing@plato.local",
			Password: "admin12345",
		})
		require.ErrorIs(t, err, domain.ErrInvalidInput)
	})
}

func TestAuthService_ChangePassword(t *testing.T) {
	ctx := context.Background()
	repo := newMockUserRepository()
	user := repo.seed("admin@plato.local", "admin12345", domain.RoleAdmin)

	cfg := DefaultAuthConfig("test-secret")
	svc := NewAuthService(repo, cfg)

	t.Run("successful change", func(t *testing.T) {
		err := svc.ChangePassword(ctx, user.ID, &dto.ChangePasswordRequest{
			OldPassword: "admin12345",
			NewPassword: "new-password-123",
		})
		require.NoError(t, err)

		_, err = svc.Login(ctx, &dto.LoginRequest{
			Email:    user.Email,
			Password: "new-password-123",
		})
		require.NoError(t, err)
	})

	t.Run("wrong old password", func(t *testing.T) {
		err := svc.ChangePassword(ctx, user.ID, &dto.ChangePasswordRequest{
			OldPassword: "wrong",
			NewPassword: "new-password-123",
		})
		require.ErrorIs(t, err, domain.ErrInvalidInput)
	})
}

func TestAuthService_Me(t *testing.T) {
	ctx := context.Background()
	repo := newMockUserRepository()
	user := repo.seed("admin@plato.local", "admin12345", domain.RoleAdmin)

	cfg := DefaultAuthConfig("test-secret")
	svc := NewAuthService(repo, cfg)

	resp, err := svc.Me(ctx, user.ID)
	require.NoError(t, err)
	require.Equal(t, user.ID, resp.ID)
	require.Equal(t, user.Email, resp.Email)
	require.Equal(t, user.Role, resp.Role)

	_, err = svc.Me(ctx, "missing-id")
	require.ErrorIs(t, err, domain.ErrNotFound)
}

func TestAuthService_TokenValidation(t *testing.T) {
	ctx := context.Background()
	repo := newMockUserRepository()
	repo.seed("admin@plato.local", "admin12345", domain.RoleAdmin)

	cfg := DefaultAuthConfig("test-secret")
	svc := NewAuthService(repo, cfg)

	resp, err := svc.Login(ctx, &dto.LoginRequest{
		Email:    "admin@plato.local",
		Password: "admin12345",
	})
	require.NoError(t, err)

	t.Run("valid access token", func(t *testing.T) {
		claims, err := svc.ParseAccessToken(resp.AccessToken)
		require.NoError(t, err)
		require.Equal(t, "admin@plato.local", claims.Email)
		require.Equal(t, domain.RoleAdmin, claims.Role)
	})

	t.Run("valid refresh token", func(t *testing.T) {
		claims, err := svc.ParseRefreshToken(resp.RefreshToken)
		require.NoError(t, err)
		require.Equal(t, "user-1", claims.UserID)
	})

	t.Run("access token rejected as refresh", func(t *testing.T) {
		_, err := svc.ParseRefreshToken(resp.AccessToken)
		require.Error(t, err)
	})

	t.Run("tampered token rejected", func(t *testing.T) {
		_, err := svc.ParseAccessToken(resp.AccessToken + "tampered")
		require.Error(t, err)
	})
}
