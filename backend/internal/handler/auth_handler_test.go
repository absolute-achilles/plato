package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/dto"
	"github.com/absolute-achilles/plato/internal/service"
	"github.com/absolute-achilles/plato/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

type mockAuthService struct {
	loginFunc          func(ctx context.Context, req *dto.LoginRequest) (*dto.TokenResponse, error)
	changePasswordFunc func(ctx context.Context, userID string, req *dto.ChangePasswordRequest) error
	parseAccessToken   func(token string) (*service.AccessTokenClaims, error)
	parseRefreshToken  func(token string) (*service.RefreshTokenClaims, error)
}

func (m *mockAuthService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.TokenResponse, error) {
	return m.loginFunc(ctx, req)
}

func (m *mockAuthService) ChangePassword(ctx context.Context, userID string, req *dto.ChangePasswordRequest) error {
	return m.changePasswordFunc(ctx, userID, req)
}

func (m *mockAuthService) ParseAccessToken(token string) (*service.AccessTokenClaims, error) {
	return m.parseAccessToken(token)
}

func (m *mockAuthService) ParseRefreshToken(token string) (*service.RefreshTokenClaims, error) {
	return m.parseRefreshToken(token)
}

func TestAuthHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)

	authSvc := &mockAuthService{
		loginFunc: func(ctx context.Context, req *dto.LoginRequest) (*dto.TokenResponse, error) {
			return &dto.TokenResponse{
				AccessToken:  "access-token",
				RefreshToken: "refresh-token",
				ExpiresIn:    3600,
				User: dto.UserResponse{
					ID:    "user-1",
					Email: req.Email,
					Role:  domain.RoleAdmin,
				},
			}, nil
		},
	}

	handler := NewAuthHandler(authSvc)
	r := gin.New()
	api := r.Group("/api/v1")
	handler.RegisterRoutes(api)

	body, _ := json.Marshal(dto.LoginRequest{Email: "admin@plato.local", Password: "admin12345"})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var env response.Envelope
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &env))
	data, _ := json.Marshal(env.Data)
	var resp dto.TokenResponse
	require.NoError(t, json.Unmarshal(data, &resp))
	require.Equal(t, "access-token", resp.AccessToken)

	cookies := w.Result().Cookies()
	require.Len(t, cookies, 2)
}

func TestAuthHandler_LoginValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	authSvc := &mockAuthService{}
	handler := NewAuthHandler(authSvc)
	r := gin.New()
	api := r.Group("/api/v1")
	handler.RegisterRoutes(api)

	body, _ := json.Marshal(map[string]string{"email": "not-an-email"})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAuthHandler_ChangePassword(t *testing.T) {
	gin.SetMode(gin.TestMode)

	authSvc := &mockAuthService{
		changePasswordFunc: func(ctx context.Context, userID string, req *dto.ChangePasswordRequest) error {
			return nil
		},
		parseAccessToken: func(token string) (*service.AccessTokenClaims, error) {
			return &service.AccessTokenClaims{
				UserID: "user-1",
				Email:  "admin@plato.local",
				Role:   domain.RoleAdmin,
				Type:   "access",
			}, nil
		},
	}

	handler := NewAuthHandler(authSvc)
	r := gin.New()
	api := r.Group("/api/v1")
	handler.RegisterRoutes(api)

	body, _ := json.Marshal(dto.ChangePasswordRequest{OldPassword: "old", NewPassword: "new123456"})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/change-password", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid-token")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusNoContent, w.Code)
}
