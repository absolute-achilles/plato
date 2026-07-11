package dto

import (
	"time"

	"github.com/absolute-achilles/plato/internal/domain"
)

// CreateParentRequest is the HTTP input for creating a parent.
type CreateParentRequest struct {
	Username    string                 `json:"username"     binding:"required"`
	Email       string                 `json:"email"        binding:"required,email"`
	Password    string                 `json:"password"     binding:"required,min=8"`
	PhoneNumber string                 `json:"phone_number,omitempty"`
	Type        domain.ParentRelationship `json:"type"         binding:"required,oneof=father mother guardian other"`
	StudentIDs  []string               `json:"student_ids"  binding:"required,min=1,dive,uuid"`
}

// ParentResponse is the parent shape returned to the client.
type ParentResponse struct {
	ID          string                    `json:"id"`
	Username    string                    `json:"username"`
	Name        string                    `json:"name"`
	Email       string                    `json:"email"`
	Role        domain.Role               `json:"role"`
	PhoneNumber *string                   `json:"phone_number,omitempty"`
	Type        domain.ParentRelationship `json:"type"`
	StudentIDs  []string                  `json:"student_ids"`
	CreatedAt   time.Time                 `json:"created_at"`
}
