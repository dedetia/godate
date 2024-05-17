package service

import (
	"context"
	"github.com/dedetia/godate/internal/core/domain"
)

type PurchaseService interface {
	PurchasePremium(ctx context.Context, request *domain.PurchasePremium) error
}
