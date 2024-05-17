package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/dedetia/godate/config"
	"github.com/dedetia/godate/internal/core/domain"
	"github.com/dedetia/godate/internal/core/port/registry"
	"github.com/dedetia/godate/pkg/auth"
	"github.com/dedetia/godate/pkg/response/custerr"
	"github.com/dedetia/godate/pkg/utils"
	"github.com/dedetia/godate/pkg/validator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
	"time"
)

type UserService struct {
	repository registry.RepositoryRegistry
	basePhotos string
}

func NewUserService(cfg *config.MainConfig, repository registry.RepositoryRegistry) *UserService {
	return &UserService{
		repository: repository,
		basePhotos: cfg.BaseUrlPhoto,
	}
}

func (us *UserService) Login(ctx context.Context, request *domain.LoginRequest) (*domain.LoginResponse, error) {
	err := validator.Validate(request)
	if err != nil {
		return nil, custerr.BadRequest(err)
	}

	user, err := us.repository.GetUserRepository().GetByPhoneNumber(ctx, request.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if user.CreatedAt.IsZero() {
		return nil, custerr.NotFound(errors.New("phone number not found"))
	}

	err = utils.Verify(utils.AlgArgon, request.Password, user.Password)
	if err != nil {
		return nil, custerr.Unauthorized(err)
	}

	token, err := us.generateToken(user)
	if err != nil {
		return nil, err
	}

	res := &domain.LoginResponse{
		ID:          user.ID,
		AccessToken: token,
	}

	return res, nil
}

func (us *UserService) generateToken(user *domain.User) (string, error) {
	claims := auth.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "godate",
			Subject:   user.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Name: user.Name,
	}

	return auth.GenerateToken(claims)
}

func (us *UserService) Signup(ctx context.Context, request *domain.SignupRequest) (*domain.SignupResponse, error) {
	err := validator.Validate(request)
	if err != nil {
		return nil, custerr.BadRequest(err)
	}

	user, err := us.repository.GetUserRepository().GetByPhoneNumberOrEmail(ctx, request.PhoneNumber, request.Email)
	if err != nil {
		return nil, err
	}

	if !user.CreatedAt.IsZero() {
		return nil, custerr.Conflict(errors.New("phone number or email is registered"))
	}

	passwordHash, err := utils.Hash(utils.AlgArgon, request.Password)
	if err != nil {
		return nil, err
	}

	user = &domain.User{
		ID:          ulid.Make().String(),
		PhoneNumber: request.PhoneNumber,
		Password:    passwordHash,
		Email:       request.Email,
		Name:        request.Name,
		DOB:         request.DOB,
		Age:         us.calculateAge(request.DOB),
		Gender:      request.Gender,
		Photos:      request.Photos,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err = us.repository.GetUserRepository().Create(ctx, user)
	if err != nil {
		return nil, err
	}

	res := &domain.SignupResponse{
		ID: user.ID,
	}

	return res, nil
}

func (us *UserService) calculateAge(dob string) int {
	birthDate, _ := time.Parse("2006-01-02", dob)
	now := time.Now().UTC()

	age := now.Year() - birthDate.Year()

	if now.YearDay() < birthDate.YearDay() {
		age--
	}

	return age
}

func (us *UserService) Profile(ctx context.Context, request *domain.ProfileRequest) (*domain.ProfileResponse, error) {
	id := auth.GetUserContext(ctx).Subject

	user, err := us.repository.GetUserRepository().GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	swipes, err := us.repository.GetSwipeRepository().GetUserSwipe(ctx, &domain.UserSwipe{
		UserID:    user.ID,
		CreatedAt: startOfDay,
	})
	if err != nil {
		return nil, err
	}

	if !user.CanSwipe(len(swipes)) {
		return nil, custerr.Forbidden(errors.New("daily swipe limit reached"))
	}

	userRec := &domain.UserRecommendation{
		SwipedUserIDs: []string{user.ID},
		Limit:         10,
		Skip:          int64((request.Page - 1) * 10),
	}

	swipedUserIDs := us.getSwipedUserIDs(swipes)
	userRec.SwipedUserIDs = append(userRec.SwipedUserIDs, swipedUserIDs...)

	recs, err := us.repository.GetUserRepository().GetUserRecommendation(ctx, userRec)
	if err != nil || len(recs) == 0 {
		return nil, err
	}

	totalData, err := us.repository.GetUserRepository().CountUserRecommendation(ctx, userRec)
	if err != nil {
		return nil, err
	}

	return us.profileResponse(recs, request.Page, totalData), nil
}

func (us *UserService) profileResponse(recs []*domain.User, page int, totalData int64) *domain.ProfileResponse {
	res := &domain.ProfileResponse{
		Profiles: make([]*domain.Profile, 0),
		Pagination: domain.Pagination{
			Page:      page,
			TotalData: totalData,
		},
	}

	for _, v := range recs {
		profile := &domain.Profile{
			ID:          v.ID,
			PhoneNumber: v.PhoneNumber,
			Email:       v.Email,
			Name:        v.Name,
			DOB:         v.DOB,
			Age:         v.Age,
			Gender:      v.Gender,
			Feature:     v.Feature,
		}
		for _, photo := range v.Photos {
			profile.Photos = append(profile.Photos, fmt.Sprintf("%s/%s", us.basePhotos, photo.Name))
		}

		res.Profiles = append(res.Profiles, profile)
	}

	return res
}

func (us *UserService) getSwipedUserIDs(swipes []*domain.Swipe) []string {
	ids := make([]string, 0)
	for _, swipe := range swipes {
		ids = append(ids, swipe.TargetUserID)
	}
	return ids
}
