package service

import (
	"context"
	"errors"
	"github.com/dedetia/godate/internal/core/domain"
	"github.com/dedetia/godate/internal/core/port/registry"
	"github.com/dedetia/godate/pkg/auth"
	"github.com/dedetia/godate/pkg/response/custerr"
	"github.com/dedetia/godate/pkg/validator"
	"github.com/dedetia/godate/shared/constant"
	"time"
)

type PurchaseService struct {
	repository registry.RepositoryRegistry
}

func NewPurchaseService(repository registry.RepositoryRegistry) *PurchaseService {
	return &PurchaseService{
		repository: repository,
	}
}

func (ps *PurchaseService) PurchasePremium(ctx context.Context, request *domain.PurchasePremium) error {
	id := auth.GetUserContext(ctx).Subject

	err := validator.Validate(request)
	if err != nil {
		return custerr.BadRequest(err)
	}

	user, err := ps.repository.GetUserRepository().GetByID(ctx, id)
	if err != nil {
		return err
	}

	if user.Feature.IsPremium {
		return custerr.Forbidden(errors.New("user has purchased the premium feature"))
	}

	switch constant.PremiumPackage(request.PackageType) {
	case constant.PremiumPackageVerifiedLabel:
		user.Feature.IsVerified = true
		user.Feature.PremiumFeature = constant.PremiumPackageVerifiedLabel.String()
	case constant.PremiumPackageNoSwipeQuota:
		user.Feature.PremiumFeature = constant.PremiumPackageNoSwipeQuota.String()
	}

	user.Feature.IsPremium = true
	user.UpdatedAt = time.Now().UTC()

	return ps.repository.GetUserRepository().Update(ctx, user)
}
