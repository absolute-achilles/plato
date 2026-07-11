package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/absolute-achilles/plato/internal/domain"
	"github.com/absolute-achilles/plato/internal/dto"
	"github.com/absolute-achilles/plato/internal/repository"
	"github.com/absolute-achilles/plato/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

// AuthConfig configures JWT signing.
type AuthConfig struct {
	Secret         string
	AccessTokenTTL time.Duration
	RefreshTokenTTL time.Duration
}

// DefaultAuthConfig returns sensible defaults for local development.
func DefaultAuthConfig(secret string) AuthConfig {
	return AuthConfig{
		Secret:          secret,
		AccessTokenTTL:  1 * time.Hour,
		RefreshTokenTTL: 7 * 24 * time.Hour,
	}
}

type AuthService interface {
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.TokenResponse, error)
	Me(ctx context.Context, userID string) (*dto.UserResponse, error)
	ChangePassword(ctx context.Context, userID string, req *dto.ChangePasswordRequest) error
	ParseAccessToken(tokenString string) (*AccessTokenClaims, error)
	ParseRefreshToken(tokenString string) (*RefreshTokenClaims, error)
}

type authService struct {
	userRepo repository.UserRepository
	cfg      AuthConfig
}

type AccessTokenClaims struct {
	UserID string      `json:"user_id"`
	Email  string      `json:"email"`
	Role   domain.Role `json:"role"`
	Type   string      `json:"type"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}

func NewAuthService(userRepo repository.UserRepository, cfg AuthConfig) AuthService {
	return &authService{userRepo: userRepo, cfg: cfg}
}

func (s *authService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.TokenResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.ErrInvalidInput
		}
		return nil, fmt.Errorf("authService.Login: %w", err)
	}

	if !utils.CheckHashPassword(req.Password, user.HashPassword) {
		return nil, domain.ErrInvalidInput
	}

	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		return nil, fmt.Errorf("authService.Login: %w", err)
	}

	refreshToken, err := s.generateRefreshToken(user)
	if err != nil {
		return nil, fmt.Errorf("authService.Login: %w", err)
	}

	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int(s.cfg.AccessTokenTTL.Seconds()),
		User:         toUserResponse(user),
	}, nil
}

func (s *authService) Me(ctx context.Context, userID string) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("authService.Me: %w", err)
	}
	resp := toUserResponse(user)
	return &resp, nil
}

func toUserResponse(user *domain.User) dto.UserResponse {
	return dto.UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		Name:        user.Username,
		Email:       user.Email,
		Role:        user.Role,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
	}
}

func (s *authService) ChangePassword(ctx context.Context, userID string, req *dto.ChangePasswordRequest) error {
	if err := s.userRepo.ChangePassword(ctx, userID, req.OldPassword, req.NewPassword); err != nil {
		if errors.Is(err, repository.ErrPasswordDoesNotMatch) {
			return domain.ErrInvalidInput
		}
		return fmt.Errorf("authService.ChangePassword: %w", err)
	}
	return nil
}

func (s *authService) ParseAccessToken(tokenString string) (*AccessTokenClaims, error) {
	claims := &AccessTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.cfg.Secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	if !token.Valid {
		return nil, domain.ErrInvalidInput
	}
	if claims.Type != "access" {
		return nil, domain.ErrInvalidInput
	}
	return claims, nil
}

func (s *authService) ParseRefreshToken(tokenString string) (*RefreshTokenClaims, error) {
	claims := &RefreshTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.cfg.Secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	if !token.Valid {
		return nil, domain.ErrInvalidInput
	}
	if claims.Type != "refresh" {
		return nil, domain.ErrInvalidInput
	}
	return claims, nil
}

func (s *authService) generateAccessToken(user *domain.User) (string, error) {
	now := time.Now()
	claims := AccessTokenClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.cfg.AccessTokenTTL)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.cfg.Secret))
}

func (s *authService) generateRefreshToken(user *domain.User) (string, error) {
	now := time.Now()
	claims := RefreshTokenClaims{
		UserID: user.ID,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.cfg.RefreshTokenTTL)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.cfg.Secret))
}
