package dto

import (
	"time"

	"github.com/absolute-achilles/plato/internal/domain"
)

// CreateStudentRequest is the HTTP input for creating a student.
type CreateStudentRequest struct {
	Username    string         `json:"username"     binding:"required"`
	Email       string         `json:"email"        binding:"required,email"`
	Password    string         `json:"password"     binding:"required,min=8"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	GradeLevel  domain.GradeLevel `json:"grade_level" binding:"required,oneof=Grade 1 Grade 2 Grade 3 Grade 4 Grade 5 Grade 6 Grade 7 Grade 8 Grade 9 Grade 10 Grade 11 Grade 12"`
}

// StudentResponse is the student shape returned to the client.
type StudentResponse struct {
	ID          string         `json:"id"`
	Username    string         `json:"username"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Role        domain.Role    `json:"role"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	GradeLevel  domain.GradeLevel `json:"grade_level"`
	CreatedAt   time.Time      `json:"created_at"`
}
