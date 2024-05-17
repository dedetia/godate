package service

import (
	"context"
	"github.com/dedetia/godate/internal/core/domain"
)

type UserService interface {
	Login(ctx context.Context, request *domain.LoginRequest) (*domain.LoginResponse, error)
	Signup(ctx context.Context, request *domain.SignupRequest) (*domain.SignupResponse, error)
	Profile(ctx context.Context, request *domain.ProfileRequest) (*domain.ProfileResponse, error)
}
