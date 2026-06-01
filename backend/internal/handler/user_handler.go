package handler

import (
	"strconv"

	"github.com/absolute-achilles/plato/internal/dto"
	"github.com/absolute-achilles/plato/internal/service"
	"github.com/absolute-achilles/plato/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.UserService // ✅ interface, not concrete struct
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) RegisterRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	users.POST("", h.createUser)
	users.GET("/:id", h.getUser)
}

func (h *UserHandler) createUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "VALIDATION_ERROR", err.Error())
		return
	}

	user, err := h.svc.CreateUser(c.Request.Context(), &req)
	if err != nil {
		response.InternalServerError(c)
		return
	}

	response.Created(c, dto.ToUserResponse(user))
}

func (h *UserHandler) getUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "INVALID_ID", "id must be a number")
		return
	}

	user, err := h.svc.GetUser(c.Request.Context(), id)
	if err != nil {
		response.NotFound(c, "user not found")
		return
	}

	response.OK(c, dto.ToUserResponse(user))
}
