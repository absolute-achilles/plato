package handler

import (
	"time"

	"github.com/absolute-achilles/plato/internal/dto"
	"github.com/absolute-achilles/plato/internal/middleware"
	"github.com/absolute-achilles/plato/internal/service"
	"github.com/absolute-achilles/plato/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc service.AuthService
}

func NewAuthHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) RegisterRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/change-password", middleware.AuthMiddleware(h.svc), h.changePassword)
		auth.GET("/me", middleware.AuthMiddleware(h.svc), h.me)
	}
}

func (h *AuthHandler) login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "VALIDATION_ERROR", err.Error())
		return
	}

	tokens, err := h.svc.Login(c.Request.Context(), &req)
	if err != nil {
		response.Unauthorized(c, "invalid email or password")
		return
	}

	setTokenCookies(c, tokens.AccessToken, tokens.RefreshToken)
	response.OK(c, tokens)
}

func (h *AuthHandler) changePassword(c *gin.Context) {
	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "VALIDATION_ERROR", err.Error())
		return
	}

	userID := middleware.GetUserID(c)
	if userID == "" {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	if err := h.svc.ChangePassword(c.Request.Context(), userID, &req); err != nil {
		response.BadRequest(c, "INVALID_PASSWORD", err.Error())
		return
	}

	response.NoContent(c)
}

func (h *AuthHandler) me(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == "" {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	user, err := h.svc.Me(c.Request.Context(), userID)
	if err != nil {
		response.Unauthorized(c, "user not found")
		return
	}

	response.OK(c, user)
}

func setTokenCookies(c *gin.Context, accessToken, refreshToken string) {
	c.SetCookie("access_token", accessToken, int((1*time.Hour).Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, int((7*24*time.Hour).Seconds()), "/", "", false, true)
}

func clearTokenCookies(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
}
