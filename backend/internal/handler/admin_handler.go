package handler

import (
	"errors"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/dto"
	"github.com/absolute-achilles/plato/internal/middleware"
	"github.com/absolute-achilles/plato/internal/service"
	"github.com/absolute-achilles/plato/pkg/response"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	svc     service.AdminService
	authSvc service.AuthService
}

func NewAdminHandler(svc service.AdminService, authSvc service.AuthService) *AdminHandler {
	return &AdminHandler{svc: svc, authSvc: authSvc}
}

func (h *AdminHandler) RegisterRoutes(api *gin.RouterGroup) {
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(h.authSvc))
	admin.Use(middleware.AdminRoleMiddleware())
	{
		admin.POST("/teachers", h.createTeacher)
		admin.POST("/students", h.createStudent)
		admin.POST("/parents", h.createParent)
	}
}

func (h *AdminHandler) createTeacher(c *gin.Context) {
	var req dto.CreateTeacherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "VALIDATION_ERROR", err.Error())
		return
	}

	teacher, err := h.svc.CreateTeacher(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicate) {
			response.Conflict(c, "DUPLICATE_USER", "username or email already exists")
			return
		}
		response.InternalServerError(c)
		return
	}

	response.Created(c, teacher)
}

func (h *AdminHandler) createStudent(c *gin.Context) {
	var req dto.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "VALIDATION_ERROR", err.Error())
		return
	}

	student, err := h.svc.CreateStudent(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicate) {
			response.Conflict(c, "DUPLICATE_USER", "username or email already exists")
			return
		}
		response.InternalServerError(c)
		return
	}

	response.Created(c, student)
}

func (h *AdminHandler) createParent(c *gin.Context) {
	var req dto.CreateParentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "VALIDATION_ERROR", err.Error())
		return
	}

	parent, err := h.svc.CreateParent(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicate) {
			response.Conflict(c, "DUPLICATE_USER", "username or email already exists")
			return
		}
		response.InternalServerError(c)
		return
	}

	response.Created(c, parent)
}
