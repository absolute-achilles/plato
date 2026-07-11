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

type mockAdminService struct {
	createTeacher func(ctx context.Context, req *dto.CreateTeacherRequest) (*dto.TeacherResponse, error)
	createStudent func(ctx context.Context, req *dto.CreateStudentRequest) (*dto.StudentResponse, error)
	createParent  func(ctx context.Context, req *dto.CreateParentRequest) (*dto.ParentResponse, error)
}

func (m *mockAdminService) CreateTeacher(ctx context.Context, req *dto.CreateTeacherRequest) (*dto.TeacherResponse, error) {
	return m.createTeacher(ctx, req)
}

func (m *mockAdminService) CreateStudent(ctx context.Context, req *dto.CreateStudentRequest) (*dto.StudentResponse, error) {
	return m.createStudent(ctx, req)
}

func (m *mockAdminService) CreateParent(ctx context.Context, req *dto.CreateParentRequest) (*dto.ParentResponse, error) {
	return m.createParent(ctx, req)
}

func setupAdminHandler() (*gin.Engine, *mockAdminService, *mockAuthService) {
	gin.SetMode(gin.TestMode)

	adminSvc := &mockAdminService{
		createTeacher: func(ctx context.Context, req *dto.CreateTeacherRequest) (*dto.TeacherResponse, error) {
			return &dto.TeacherResponse{
				ID:         "teacher-1",
				Username:   req.Username,
				Email:      req.Email,
				Role:       domain.RoleTeacher,
				Department: req.Department,
			}, nil
		},
	}

	authSvc := &mockAuthService{
		parseAccessToken: func(token string) (*service.AccessTokenClaims, error) {
			return &service.AccessTokenClaims{
				UserID: "user-1",
				Email:  "admin@plato.local",
				Role:   domain.RoleAdmin,
				Type:   "access",
			}, nil
		},
	}

	handler := NewAdminHandler(adminSvc, authSvc)
	r := gin.New()
	api := r.Group("/api/v1")
	handler.RegisterRoutes(api)

	return r, adminSvc, authSvc
}

func TestAdminHandler_CreateTeacher(t *testing.T) {
	r, _, _ := setupAdminHandler()

	body, _ := json.Marshal(dto.CreateTeacherRequest{
		Username:   "budi.santoso",
		Email:      "budi@plato.edu",
		Password:   "password123",
		Department: domain.DepartmentMathematics,
	})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/teachers", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid-admin-token")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	var env response.Envelope
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &env))
	data, _ := json.Marshal(env.Data)
	var resp dto.TeacherResponse
	require.NoError(t, json.Unmarshal(data, &resp))
	require.Equal(t, "budi.santoso", resp.Username)
	require.Equal(t, domain.DepartmentMathematics, resp.Department)
}

func TestAdminHandler_CreateTeacherForbiddenForNonAdmin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	adminSvc := &mockAdminService{}
	authSvc := &mockAuthService{
		parseAccessToken: func(token string) (*service.AccessTokenClaims, error) {
			return &service.AccessTokenClaims{
				UserID: "user-1",
				Email:  "teacher@plato.edu",
				Role:   domain.RoleTeacher,
				Type:   "access",
			}, nil
		},
	}

	handler := NewAdminHandler(adminSvc, authSvc)
	r := gin.New()
	api := r.Group("/api/v1")
	handler.RegisterRoutes(api)

	body, _ := json.Marshal(dto.CreateTeacherRequest{
		Username:   "budi.santoso",
		Email:      "budi@plato.edu",
		Password:   "password123",
		Department: domain.DepartmentMathematics,
	})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/teachers", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid-teacher-token")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusForbidden, w.Code)
}

func TestAdminHandler_CreateTeacherValidationError(t *testing.T) {
	r, _, _ := setupAdminHandler()

	body, _ := json.Marshal(map[string]string{"email": "not-an-email"})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/teachers", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer valid-admin-token")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAdminHandler_CreateTeacherUnauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	adminSvc := &mockAdminService{}
	authSvc := &mockAuthService{}

	handler := NewAdminHandler(adminSvc, authSvc)
	r := gin.New()
	api := r.Group("/api/v1")
	handler.RegisterRoutes(api)

	body, _ := json.Marshal(dto.CreateTeacherRequest{
		Username:   "budi.santoso",
		Email:      "budi@plato.edu",
		Password:   "password123",
		Department: domain.DepartmentMathematics,
	})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/admin/teachers", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusUnauthorized, w.Code)
}
