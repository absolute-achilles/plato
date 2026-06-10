package dto

import "github.com/absolute-achilles/plato/internal/domain"

// CreateTeacherRequest — HTTP input
type CreateTeacherRequest struct {
	Name     string `json:"name"  binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	// Role     domain.Role `json:"role"  binding:"required,oneof=student teacher parent admin"`
}

// UserResponse — controls exactly what fields go back to the client
type TeacherResponse struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Role  domain.Role `json:"role"`
}
