package dto

import (
	"time"

	"github.com/absolute-achilles/plato/internal/domain"
)

// CreateTeacherRequest is the HTTP input for creating a teacher.
type CreateTeacherRequest struct {
	Username    string           `json:"username"     binding:"required"`
	Email       string           `json:"email"        binding:"required,email"`
	Password    string           `json:"password"     binding:"required,min=8"`
	PhoneNumber string           `json:"phone_number,omitempty"`
	Department  domain.Department `json:"department"  binding:"required,oneof=Mathematics Science English History Arts Physical Education Computer Science Other"`
}

// TeacherResponse is the teacher shape returned to the client.
type TeacherResponse struct {
	ID          string           `json:"id"`
	Username    string           `json:"username"`
	Name        string           `json:"name"`
	Email       string           `json:"email"`
	Role        domain.Role      `json:"role"`
	PhoneNumber *string          `json:"phone_number,omitempty"`
	Department  domain.Department `json:"department"`
	CreatedAt   time.Time        `json:"created_at"`
}
