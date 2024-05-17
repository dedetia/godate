package repository

import (
	"context"
	"github.com/dedetia/godate/internal/core/domain"
)

type UserRepository interface {
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*domain.User, error)
	GetByPhoneNumberOrEmail(ctx context.Context, phoneNumber, email string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetUserRecommendation(ctx context.Context, request *domain.UserRecommendation) ([]*domain.User, error)
	CountUserRecommendation(ctx context.Context, request *domain.UserRecommendation) (int64, error)
	Update(ctx context.Context, user *domain.User) error
}
