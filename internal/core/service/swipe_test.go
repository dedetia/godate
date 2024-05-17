package service

import (
	"context"
	"errors"
	"github.com/dedetia/godate/internal/core/domain"
	"github.com/dedetia/godate/pkg/auth"
	"github.com/dedetia/godate/shared/mock/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestSwipeAction(t *testing.T) {
	Convey("swipe action", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			repositoryRegistry = repository.NewMockRepositoryRegistry(ctrl)
			userRepo           = repository.NewMockUserRepository(ctrl)
			swipeRepo          = repository.NewMockSwipeRepository(ctrl)
			server             = NewSwipeService(repositoryRegistry)

			request = &domain.SwipeRequest{
				TargetUserID: "234",
				Action:       "like",
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

			userTarget = &domain.User{
				ID:          "234",
				PhoneNumber: "+6281294309911",
				Feature: domain.Feature{
					IsPremium:      false,
					IsVerified:     false,
					PremiumFeature: "",
				},
				CreatedAt: time.Now(),
			}

			swipes = []*domain.Swipe{
				{
					ID:           "123",
					UserID:       "123",
					TargetUserID: "456",
					Action:       "pass",
				},
			}

			swipesLimit = []*domain.Swipe{
				{
					ID:           "123",
					UserID:       "123",
					TargetUserID: "456",
					Action:       "pass",
				},
				{
					ID:           "123",
					UserID:       "123",
					TargetUserID: "452",
					Action:       "pass",
				},
				{
					ID:           "123",
					UserID:       "123",
					TargetUserID: "111",
					Action:       "pass",
				},
				{
					ID:           "123",
					UserID:       "123",
					TargetUserID: "222",
					Action:       "pass",
				},
				{
					ID:           "123",
					UserID:       "123",
					TargetUserID: "333",
					Action:       "pass",
				},
				{
					ID:           "123",
					UserID:       "123",
					TargetUserID: "845",
					Action:       "pass",
				},
				{
					ID:           "123",
					UserID:       "123",
					TargetUserID: "743",
					Action:       "pass",
				},
				{
					ID:           "123",
					UserID:       "123",
					TargetUserID: "363",
					Action:       "pass",
				},
				{
					ID:           "123",
					UserID:       "123",
					TargetUserID: "643",
					Action:       "pass",
				},
				{
					ID:           "123",
					UserID:       "123",
					TargetUserID: "334",
					Action:       "like",
				},
			}

			claimJwt = &auth.CustomClaims{
				RegisteredClaims: jwt.RegisteredClaims{
					Subject: "123",
				},
				Name: "test",
			}

			ctx = context.WithValue(context.Background(), auth.UserKey, claimJwt)
			err = errors.New("error")
		)

		Convey("error validator", func() {
			request.Action = "aaa"
			err := server.SwipeAction(ctx, request)
			So(err, ShouldNotBeNil)
		})

		Convey("error cannot swipe yourself", func() {
			request.TargetUserID = "123"
			err := server.SwipeAction(ctx, request)
			So(err, ShouldNotBeNil)
		})

		Convey("error get user by id", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo)
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, err)
			err := server.SwipeAction(ctx, request)
			So(err, ShouldNotBeNil)
		})

		Convey("error get target user by id", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, err)
			err := server.SwipeAction(ctx, request)
			So(err, ShouldNotBeNil)
		})

		Convey("error target user not found", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(&domain.User{}, nil)
			err := server.SwipeAction(ctx, request)
			So(err, ShouldNotBeNil)
		})

		Convey("error get user swipes", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(userTarget, nil)
			repositoryRegistry.EXPECT().GetSwipeRepository().Return(swipeRepo)
			swipeRepo.EXPECT().GetUserSwipe(gomock.Any(), gomock.Any()).Return(nil, err)

			err := server.SwipeAction(ctx, request)
			So(err, ShouldNotBeNil)
		})

		Convey("error swipe limit reached", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(userTarget, nil)
			repositoryRegistry.EXPECT().GetSwipeRepository().Return(swipeRepo).AnyTimes()
			swipeRepo.EXPECT().GetUserSwipe(gomock.Any(), gomock.Any()).Return(swipesLimit, nil)

			err := server.SwipeAction(ctx, request)
			So(err, ShouldNotBeNil)
		})

		Convey("error user target swiped", func() {
			swiped := []*domain.Swipe{
				{
					ID:           "123",
					UserID:       "123",
					TargetUserID: "234",
					Action:       "pass",
				},
			}
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(userTarget, nil)
			repositoryRegistry.EXPECT().GetSwipeRepository().Return(swipeRepo)
			swipeRepo.EXPECT().GetUserSwipe(gomock.Any(), gomock.Any()).Return(swiped, nil)

			err := server.SwipeAction(ctx, request)
			So(err, ShouldNotBeNil)
		})

		Convey("error create swipe", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(userTarget, nil)
			repositoryRegistry.EXPECT().GetSwipeRepository().Return(swipeRepo).AnyTimes()
			swipeRepo.EXPECT().GetUserSwipe(gomock.Any(), gomock.Any()).Return(swipes, nil)
			swipeRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(err)

			err := server.SwipeAction(ctx, request)
			So(err, ShouldNotBeNil)
		})

		Convey("success swipe", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(userTarget, nil)
			repositoryRegistry.EXPECT().GetSwipeRepository().Return(swipeRepo).AnyTimes()
			swipeRepo.EXPECT().GetUserSwipe(gomock.Any(), gomock.Any()).Return(swipes, nil)
			swipeRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

			err := server.SwipeAction(ctx, request)
			So(err, ShouldBeNil)
		})
	})
}
