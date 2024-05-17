package service

import (
	"context"
	"github.com/dedetia/godate/internal/core/domain"
)

type SwipeService interface {
	SwipeAction(ctx context.Context, request *domain.SwipeRequest) error
}
