package service

import (
	"context"
	"errors"
	"github.com/dedetia/godate/config"
	"github.com/dedetia/godate/internal/core/domain"
	"github.com/dedetia/godate/pkg/auth"
	"github.com/dedetia/godate/pkg/utils"
	"github.com/dedetia/godate/shared/constant"
	"github.com/dedetia/godate/shared/mock/repository"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	Convey("login", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			cfg = &config.MainConfig{
				BaseUrlPhoto: "localhost",
			}
			repositoryRegistry = repository.NewMockRepositoryRegistry(ctrl)
			userRepo           = repository.NewMockUserRepository(ctrl)
			server             = NewUserService(cfg, repositoryRegistry)

			request = &domain.LoginRequest{
				PhoneNumber: "+6281294309919",
				Password:    "G0d@t1ing",
			}

			passwordHash, _ = utils.Hash(utils.AlgArgon, "G0d@t1ing")

			user = &domain.User{
				ID:          "123",
				Password:    passwordHash,
				PhoneNumber: "+6281294309919",
				Feature: domain.Feature{
					IsPremium:      false,
					IsVerified:     false,
					PremiumFeature: "",
				},
				CreatedAt: time.Now(),
			}

			ctx = context.Background()
			err = errors.New("error")
		)

		_ = auth.Configure("MIIEpgIBAAKCAQEA1WdEXZHnpEhiDAzWEN1X0dwEUtyNHqzxaq5hP9S33Lgtb9AQMCh7zMU98eG48LrjMIEnhxiS4HDgoi5pOAZVlZcKxGBmvDc0jxizfUl17BND9fJF6yB2DE50fyIXL+rZpHycV3cqpkm3PS09THYX2LvZmpb4n6Wpcp5qIyiGr7WG7o2HfZNilOVJwqmOT42dZ4cuxX2EW4VFNONSxTMO2/X0M7K92WRcXv9KD/va4trofAiv2+O2kbppim8Qpi2tiAqyi/Pr9LiqRZwa19oO5hDr+klfOj2yFI+2eW7StkAtLVJUPcIx6E9suy6zG0LwITfc2shawqwtsKKEwpIK1QIDAQABAoIBAQCGIj+VdMUdvKVsH5FZzlaJwPoyvxAwjNG9lVfpECJ1KIrese/K5VdTUVLrO07MeRud/EBFKQwA6NI4/mUCYvDecq7A2jsY6LYvj34aLNdjCIT6DUsnTCMG/zU4R8w9QSeFvRFj5LI5DTKQ0GOsMLoyb3iKM4SYjD8inTHnYWyu+YuaYrY5JhvNKbCoDqpYWhMCOUR0dw+YyzhqmjwaKg3o1u4MTVBxh1qCpqLVWZ6sMOPQbDMmlS/v66rFrNVvmCqbDWwwEphc+BnWM9XPY+g3Vn3/ZsKoTSNlfWI1F11ZWubkfvt6LwaxlMTAjdq4U3Y6V8dASeOx0Gxa8kG/hGBBAoGBAOKIDgnZMtTMJfl8jf1ILsQ+dNwK3/VcJk3Z7ZapoyWg9ukBQVKInDL1+ov9ZByT0jj7EBmW6o5mmc6EzsQxCjG0QO+DGitrTpWR5ifvssbs3WkqMgbKF5Tz0MJ/9D33eN9iFEjVqyJXjaLA9HdPVSLPKJevYmVWDFaPMDEP8LMFAoGBAPEqBXTH53yRYlmKna+yrk549YuNGJ7jqTYZeNwHoxOaQvh1XDEJGnfjhue/hmK/3qsVTvLJymagAKOqF9MFS8ppAGyl29ss6zwhziYLuLzITmGL59Pw/5p9+rvNt8rmwxON0UtqnC9/rrbSaA1vJ8/07uAxiZf1NYzYO/aN1SGRAoGBAKfYpYY4j8hKZ0y/NDnaNQSlPlMYH68eEyeV9Muwb7je1nP4wRzVKd88kOMO4hGmmZostFYxkyPl88qobsfBiksfwwl0e3x2aui6DO3EVhO8x6U3ZY/QR77PFPw4cJFFfyMM+fipkL7GXqScEcchWfSLyAj0I5TwN/4e5FdF91O9AoGBAMgxbNQTeesjOLRB6EJInn+P061jlDOZowbAwF5OjKYiIUPlEIG4H9uz6XIJwEHLKsl0Z9QNhNIKMl2qPhqzQ8YjwfFvAYIA2MlS+rEEe/dihAZfwDNk1JnnyDMMQ2zQgNGDoWDsf/jCEkO7iBrW0gLEPWOoW6LkL+7aNXSnKmyxAoGBAJ2kMwIjbqxtxwEaNClcQ9cUDWxREgVowaV5IP1Ae21JonaaYkpzkQsiEOy1hciaK3uoqqLp6CFLLM8RKzrRh+qb1Xu5/FwNePaxrQ/ckDAN1AwWtL50Rt6J1HQn+sjRIgNc7weUhqI9MWHhSkm0qYXv4/HfU2Or8KUH0RJDf+nf")

		Convey("error validator", func() {
			request.PhoneNumber = ""
			res, err := server.Login(ctx, request)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("error get user by phone number", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo)
			userRepo.EXPECT().GetByPhoneNumber(gomock.Any(), gomock.Any()).Return(nil, err)
			res, err := server.Login(ctx, request)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("error phone number not found", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo)
			userRepo.EXPECT().GetByPhoneNumber(gomock.Any(), gomock.Any()).Return(&domain.User{}, nil)
			res, err := server.Login(ctx, request)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("error invalid password", func() {
			request.Password = "G0go012345"
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo)
			userRepo.EXPECT().GetByPhoneNumber(gomock.Any(), gomock.Any()).Return(user, nil)
			res, err := server.Login(ctx, request)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("login success", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo)
			userRepo.EXPECT().GetByPhoneNumber(gomock.Any(), gomock.Any()).Return(user, nil)
			res, err := server.Login(ctx, request)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
		})
	})
}

func TestSignup(t *testing.T) {
	Convey("signup", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			cfg = &config.MainConfig{
				BaseUrlPhoto: "localhost",
			}
			repositoryRegistry = repository.NewMockRepositoryRegistry(ctrl)
			userRepo           = repository.NewMockUserRepository(ctrl)
			server             = NewUserService(cfg, repositoryRegistry)

			request = &domain.SignupRequest{
				PhoneNumber: "+6281294309919",
				Password:    "G0d@t1ing",
				Email:       "dede@gmail.com",
				Name:        "Dede",
				DOB:         "2024-05-17",
				Gender:      "Male",
				Photos: []*domain.File{
					{
						Name: "test.jpg",
						Path: "/assets/test.jpg",
					},
				},
			}

			passwordHash, _ = utils.Hash(utils.AlgArgon, "G0d@t1ing")
			user            = &domain.User{
				ID:          "123",
				Password:    passwordHash,
				PhoneNumber: "+6281294309919",
				Feature: domain.Feature{
					IsPremium:      false,
					IsVerified:     false,
					PremiumFeature: "",
				},
				CreatedAt: time.Now(),
			}

			ctx = context.Background()
			err = errors.New("error")
		)

		_ = auth.Configure("MIIEpgIBAAKCAQEA1WdEXZHnpEhiDAzWEN1X0dwEUtyNHqzxaq5hP9S33Lgtb9AQMCh7zMU98eG48LrjMIEnhxiS4HDgoi5pOAZVlZcKxGBmvDc0jxizfUl17BND9fJF6yB2DE50fyIXL+rZpHycV3cqpkm3PS09THYX2LvZmpb4n6Wpcp5qIyiGr7WG7o2HfZNilOVJwqmOT42dZ4cuxX2EW4VFNONSxTMO2/X0M7K92WRcXv9KD/va4trofAiv2+O2kbppim8Qpi2tiAqyi/Pr9LiqRZwa19oO5hDr+klfOj2yFI+2eW7StkAtLVJUPcIx6E9suy6zG0LwITfc2shawqwtsKKEwpIK1QIDAQABAoIBAQCGIj+VdMUdvKVsH5FZzlaJwPoyvxAwjNG9lVfpECJ1KIrese/K5VdTUVLrO07MeRud/EBFKQwA6NI4/mUCYvDecq7A2jsY6LYvj34aLNdjCIT6DUsnTCMG/zU4R8w9QSeFvRFj5LI5DTKQ0GOsMLoyb3iKM4SYjD8inTHnYWyu+YuaYrY5JhvNKbCoDqpYWhMCOUR0dw+YyzhqmjwaKg3o1u4MTVBxh1qCpqLVWZ6sMOPQbDMmlS/v66rFrNVvmCqbDWwwEphc+BnWM9XPY+g3Vn3/ZsKoTSNlfWI1F11ZWubkfvt6LwaxlMTAjdq4U3Y6V8dASeOx0Gxa8kG/hGBBAoGBAOKIDgnZMtTMJfl8jf1ILsQ+dNwK3/VcJk3Z7ZapoyWg9ukBQVKInDL1+ov9ZByT0jj7EBmW6o5mmc6EzsQxCjG0QO+DGitrTpWR5ifvssbs3WkqMgbKF5Tz0MJ/9D33eN9iFEjVqyJXjaLA9HdPVSLPKJevYmVWDFaPMDEP8LMFAoGBAPEqBXTH53yRYlmKna+yrk549YuNGJ7jqTYZeNwHoxOaQvh1XDEJGnfjhue/hmK/3qsVTvLJymagAKOqF9MFS8ppAGyl29ss6zwhziYLuLzITmGL59Pw/5p9+rvNt8rmwxON0UtqnC9/rrbSaA1vJ8/07uAxiZf1NYzYO/aN1SGRAoGBAKfYpYY4j8hKZ0y/NDnaNQSlPlMYH68eEyeV9Muwb7je1nP4wRzVKd88kOMO4hGmmZostFYxkyPl88qobsfBiksfwwl0e3x2aui6DO3EVhO8x6U3ZY/QR77PFPw4cJFFfyMM+fipkL7GXqScEcchWfSLyAj0I5TwN/4e5FdF91O9AoGBAMgxbNQTeesjOLRB6EJInn+P061jlDOZowbAwF5OjKYiIUPlEIG4H9uz6XIJwEHLKsl0Z9QNhNIKMl2qPhqzQ8YjwfFvAYIA2MlS+rEEe/dihAZfwDNk1JnnyDMMQ2zQgNGDoWDsf/jCEkO7iBrW0gLEPWOoW6LkL+7aNXSnKmyxAoGBAJ2kMwIjbqxtxwEaNClcQ9cUDWxREgVowaV5IP1Ae21JonaaYkpzkQsiEOy1hciaK3uoqqLp6CFLLM8RKzrRh+qb1Xu5/FwNePaxrQ/ckDAN1AwWtL50Rt6J1HQn+sjRIgNc7weUhqI9MWHhSkm0qYXv4/HfU2Or8KUH0RJDf+nf")

		Convey("error validator", func() {
			request.PhoneNumber = ""
			res, err := server.Signup(ctx, request)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("error get user by phone number or email", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo)
			userRepo.EXPECT().GetByPhoneNumberOrEmail(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, err)
			res, err := server.Signup(ctx, request)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("error phone number or email is registered", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo)
			userRepo.EXPECT().GetByPhoneNumberOrEmail(gomock.Any(), gomock.Any(), gomock.Any()).Return(user, nil)
			res, err := server.Signup(ctx, request)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("error create user", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByPhoneNumberOrEmail(gomock.Any(), gomock.Any(), gomock.Any()).Return(&domain.User{}, nil)
			userRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(err)
			res, err := server.Signup(ctx, request)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("signup success", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByPhoneNumberOrEmail(gomock.Any(), gomock.Any(), gomock.Any()).Return(&domain.User{}, nil)
			userRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			res, err := server.Signup(ctx, request)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
		})
	})
}

func TestProfile(t *testing.T) {
	Convey("profile", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			cfg = &config.MainConfig{
				BaseUrlPhoto: "localhost",
			}
			repositoryRegistry = repository.NewMockRepositoryRegistry(ctrl)
			userRepo           = repository.NewMockUserRepository(ctrl)
			swipeRepo          = repository.NewMockSwipeRepository(ctrl)
			server             = NewUserService(cfg, repositoryRegistry)

			request = &domain.ProfileRequest{
				Page: 1,
			}

			user = &domain.User{
				ID:          "123",
				PhoneNumber: "+6281294309919",
				Email:       "mamang@gmail.com",
				Name:        "Mamang",
				DOB:         "1920-02-28",
				Age:         104,
				Gender:      "Male",
				Feature: domain.Feature{
					IsPremium:      true,
					IsVerified:     true,
					PremiumFeature: constant.PremiumPackageVerifiedLabel.String(),
				},
				CreatedAt: time.Now(),
				UpdatedAt: time.Time{},
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

			users = []*domain.User{
				{
					ID:          "432",
					PhoneNumber: "+6281283839999",
					Email:       "nona@gmail.com",
					Name:        "Nona",
					DOB:         "1945-08-17",
					Age:         79,
					Gender:      "Female",
					Photos: []*domain.File{
						{
							Name: "test.png",
							Path: "/test.png",
						},
					},
					Feature:   domain.Feature{},
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
			}

			ctx = context.Background()
			err = errors.New("error")
		)

		_ = auth.Configure("MIIEpgIBAAKCAQEA1WdEXZHnpEhiDAzWEN1X0dwEUtyNHqzxaq5hP9S33Lgtb9AQMCh7zMU98eG48LrjMIEnhxiS4HDgoi5pOAZVlZcKxGBmvDc0jxizfUl17BND9fJF6yB2DE50fyIXL+rZpHycV3cqpkm3PS09THYX2LvZmpb4n6Wpcp5qIyiGr7WG7o2HfZNilOVJwqmOT42dZ4cuxX2EW4VFNONSxTMO2/X0M7K92WRcXv9KD/va4trofAiv2+O2kbppim8Qpi2tiAqyi/Pr9LiqRZwa19oO5hDr+klfOj2yFI+2eW7StkAtLVJUPcIx6E9suy6zG0LwITfc2shawqwtsKKEwpIK1QIDAQABAoIBAQCGIj+VdMUdvKVsH5FZzlaJwPoyvxAwjNG9lVfpECJ1KIrese/K5VdTUVLrO07MeRud/EBFKQwA6NI4/mUCYvDecq7A2jsY6LYvj34aLNdjCIT6DUsnTCMG/zU4R8w9QSeFvRFj5LI5DTKQ0GOsMLoyb3iKM4SYjD8inTHnYWyu+YuaYrY5JhvNKbCoDqpYWhMCOUR0dw+YyzhqmjwaKg3o1u4MTVBxh1qCpqLVWZ6sMOPQbDMmlS/v66rFrNVvmCqbDWwwEphc+BnWM9XPY+g3Vn3/ZsKoTSNlfWI1F11ZWubkfvt6LwaxlMTAjdq4U3Y6V8dASeOx0Gxa8kG/hGBBAoGBAOKIDgnZMtTMJfl8jf1ILsQ+dNwK3/VcJk3Z7ZapoyWg9ukBQVKInDL1+ov9ZByT0jj7EBmW6o5mmc6EzsQxCjG0QO+DGitrTpWR5ifvssbs3WkqMgbKF5Tz0MJ/9D33eN9iFEjVqyJXjaLA9HdPVSLPKJevYmVWDFaPMDEP8LMFAoGBAPEqBXTH53yRYlmKna+yrk549YuNGJ7jqTYZeNwHoxOaQvh1XDEJGnfjhue/hmK/3qsVTvLJymagAKOqF9MFS8ppAGyl29ss6zwhziYLuLzITmGL59Pw/5p9+rvNt8rmwxON0UtqnC9/rrbSaA1vJ8/07uAxiZf1NYzYO/aN1SGRAoGBAKfYpYY4j8hKZ0y/NDnaNQSlPlMYH68eEyeV9Muwb7je1nP4wRzVKd88kOMO4hGmmZostFYxkyPl88qobsfBiksfwwl0e3x2aui6DO3EVhO8x6U3ZY/QR77PFPw4cJFFfyMM+fipkL7GXqScEcchWfSLyAj0I5TwN/4e5FdF91O9AoGBAMgxbNQTeesjOLRB6EJInn+P061jlDOZowbAwF5OjKYiIUPlEIG4H9uz6XIJwEHLKsl0Z9QNhNIKMl2qPhqzQ8YjwfFvAYIA2MlS+rEEe/dihAZfwDNk1JnnyDMMQ2zQgNGDoWDsf/jCEkO7iBrW0gLEPWOoW6LkL+7aNXSnKmyxAoGBAJ2kMwIjbqxtxwEaNClcQ9cUDWxREgVowaV5IP1Ae21JonaaYkpzkQsiEOy1hciaK3uoqqLp6CFLLM8RKzrRh+qb1Xu5/FwNePaxrQ/ckDAN1AwWtL50Rt6J1HQn+sjRIgNc7weUhqI9MWHhSkm0qYXv4/HfU2Or8KUH0RJDf+nf")

		Convey("error get user by id", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo)
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, err)
			res, err := server.Profile(ctx, request)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("error get user swipes", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo)
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			repositoryRegistry.EXPECT().GetSwipeRepository().Return(swipeRepo)
			swipeRepo.EXPECT().GetUserSwipe(gomock.Any(), gomock.Any()).Return(nil, err)
			res, err := server.Profile(ctx, request)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("error swap limit", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo)
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			repositoryRegistry.EXPECT().GetSwipeRepository().Return(swipeRepo)
			swipeRepo.EXPECT().GetUserSwipe(gomock.Any(), gomock.Any()).Return(swipesLimit, nil)
			res, err := server.Profile(ctx, request)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("error get user recommendations", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			repositoryRegistry.EXPECT().GetSwipeRepository().Return(swipeRepo)
			swipeRepo.EXPECT().GetUserSwipe(gomock.Any(), gomock.Any()).Return(swipes, nil)
			userRepo.EXPECT().GetUserRecommendation(gomock.Any(), gomock.Any()).Return(nil, err)
			res, err := server.Profile(ctx, request)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("error user recommendations is nil", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			repositoryRegistry.EXPECT().GetSwipeRepository().Return(swipeRepo)
			swipeRepo.EXPECT().GetUserSwipe(gomock.Any(), gomock.Any()).Return(swipes, nil)
			userRepo.EXPECT().GetUserRecommendation(gomock.Any(), gomock.Any()).Return(nil, nil)
			res, err := server.Profile(ctx, request)
			So(err, ShouldBeNil)
			So(res, ShouldBeNil)
		})

		Convey("error get count user recommendation", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			repositoryRegistry.EXPECT().GetSwipeRepository().Return(swipeRepo)
			swipeRepo.EXPECT().GetUserSwipe(gomock.Any(), gomock.Any()).Return(swipes, nil)
			userRepo.EXPECT().GetUserRecommendation(gomock.Any(), gomock.Any()).Return(users, nil)
			userRepo.EXPECT().CountUserRecommendation(gomock.Any(), gomock.Any()).Return(int64(0), err)
			res, err := server.Profile(ctx, request)
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})

		Convey("success", func() {
			repositoryRegistry.EXPECT().GetUserRepository().Return(userRepo).AnyTimes()
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(user, nil)
			repositoryRegistry.EXPECT().GetSwipeRepository().Return(swipeRepo)
			swipeRepo.EXPECT().GetUserSwipe(gomock.Any(), gomock.Any()).Return(swipes, nil)
			userRepo.EXPECT().GetUserRecommendation(gomock.Any(), gomock.Any()).Return(users, nil)
			userRepo.EXPECT().CountUserRecommendation(gomock.Any(), gomock.Any()).Return(int64(1), nil)
			res, err := server.Profile(ctx, request)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
		})

	})
}
