package dto

import "github.com/absolute-achilles/plato/internal/domain"

// CreateUserRequest — HTTP input
type CreateUserRequest struct {
	Name  string      `json:"name"  binding:"required"`
	Email string      `json:"email" binding:"required,email"`
	Role  domain.Role `json:"role"  binding:"required,oneof=student teacher guardian admin"`
}

// UserResponse — controls exactly what fields go back to the client
type UserResponse struct {
	ID    int64       `json:"id"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Role  domain.Role `json:"role"`
}

// ToUserResponse maps domain entity → response shape
func ToUserResponse(u *domain.User) UserResponse {
	return UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role,
	}
}
