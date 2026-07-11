package middleware

import (
	"strings"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/service"
	"github.com/absolute-achilles/plato/pkg/response"
	"github.com/gin-gonic/gin"
)

const (
	UserIDContextKey = "userID"
	UserRoleContextKey = "userRole"
	UserEmailContextKey = "userEmail"
)

// AuthMiddleware validates the access token in either the Authorization header or the access_token cookie.
func AuthMiddleware(authSvc service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)
		if token == "" {
			response.Unauthorized(c, "missing access token")
			c.Abort()
			return
		}

		claims, err := authSvc.ParseAccessToken(token)
		if err != nil {
			response.Unauthorized(c, "invalid or expired access token")
			c.Abort()
			return
		}

		c.Set(UserIDContextKey, claims.UserID)
		c.Set(UserRoleContextKey, string(claims.Role))
		c.Set(UserEmailContextKey, claims.Email)
		c.Next()
	}
}

// AdminRoleMiddleware ensures the authenticated user has the admin role.
func AdminRoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(UserRoleContextKey)
		if !exists || role != string(domain.RoleAdmin) {
			response.Forbidden(c, "admin access required")
			c.Abort()
			return
		}
		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
			return parts[1]
		}
	}

	if cookie, err := c.Cookie("access_token"); err == nil && cookie != "" {
		return cookie
	}

	return ""
}

// GetUserID returns the authenticated user ID from the Gin context.
func GetUserID(c *gin.Context) string {
	id, _ := c.Get(UserIDContextKey)
	if id == nil {
		return ""
	}
	return id.(string)
}
