package repository

import (
	"context"
	"github.com/dedetia/godate/internal/core/domain"
)

type SwipeRepository interface {
	GetUserSwipe(ctx context.Context, request *domain.UserSwipe) ([]*domain.Swipe, error)
	Create(ctx context.Context, request *domain.Swipe) error
}
