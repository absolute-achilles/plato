package dto

import "github.com/absolute-achilles/plato/internal/domain"

// LoginRequest is the input for signing in.
type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// ChangePasswordRequest is the input for updating the authenticated user's password.
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

// TokenResponse is returned after a successful login.
type TokenResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	ExpiresIn    int         `json:"expires_in"`
	User         UserResponse `json:"user"`
}

// UserResponse is the shared user shape returned to the client.
type UserResponse struct {
	ID          string      `json:"id"`
	Username    string      `json:"username"`
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	Role        domain.Role `json:"role"`
	PhoneNumber *string     `json:"phone_number,omitempty"`
	CreatedAt   string      `json:"created_at"`
}
