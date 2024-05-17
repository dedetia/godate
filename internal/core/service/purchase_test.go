package service

import (
	"context"
	"errors"
	"github.com/dedetia/godate/internal/core/domain"
	"github.com/dedetia/godate/shared/constant"
	"github.com/dedetia/godate/shared/mock/repository"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPurchase(t *testing.T) {
	Convey("purchase premium", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			repositoryRegistry = repository.NewMockRepositoryRegistry(ctrl)
			userRepo           = repository.NewMockUserRepository(ctrl)
			server             = NewPurchaseService(repositoryRegistry)

			request = &domain.PurchasePremium{
				PackageType: constant.PremiumPackageVerifiedLabel.String(),
			}

			user = &domain.User{
				ID:          "123",
				PhoneNumber: "+6281294309919",
				Feature: domain.Feature{
					IsPremium:      false,
					IsVerified:     false,
					PremiumFeature: "",
				},
			}

			ctx = context.Background()
			err = errors.New("error")
		)

		Convey("error validator", func() {
			request.PackageType = "paket"
			err := server.PurchasePremium(ctx, request)
			So(err, ShouldNotBeNil)
		})

		Convey("error get user by id", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo)
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, err)
			err := server.PurchasePremium(ctx, request)
			So(err, ShouldNotBeNil)
		})

		Convey("error user has purchased premium", func() {
			user.Feature.IsPremium = true
			user.Feature.IsVerified = true

			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo)
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			err := server.PurchasePremium(ctx, request)
			So(err, ShouldNotBeNil)
		})

		Convey("error update user feature", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			userRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(err)
			err := server.PurchasePremium(ctx, request)
			So(err, ShouldNotBeNil)
		})

		Convey("update success", func() {
			request.PackageType = constant.PremiumPackageNoSwipeQuota.String()
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			userRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			err := server.PurchasePremium(ctx, request)
			So(err, ShouldBeNil)
		})
	})

}
