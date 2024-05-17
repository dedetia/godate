package service

import (
	"context"
	"errors"
	"github.com/dedetia/godate/internal/core/domain"
	"github.com/dedetia/godate/internal/core/port/registry"
	"github.com/dedetia/godate/pkg/auth"
	"github.com/dedetia/godate/pkg/response/custerr"
	"github.com/dedetia/godate/pkg/validator"
	"github.com/oklog/ulid/v2"
	"time"
)

type SwipeService struct {
	repository registry.RepositoryRegistry
}

func NewSwipeService(repository registry.RepositoryRegistry) *SwipeService {
	return &SwipeService{
		repository: repository,
	}
}

func (ss *SwipeService) SwipeAction(ctx context.Context, request *domain.SwipeRequest) error {
	id := auth.GetUserContext(ctx).Subject

	err := validator.Validate(request)
	if err != nil {
		return err
	}

	if id == request.TargetUserID {
		return custerr.Forbidden(errors.New("cannot swipe on yourself"))
	}

	user, err := ss.repository.GetUserRepository().GetByID(ctx, id)
	if err != nil {
		return err
	}

	userTarget, err := ss.repository.GetUserRepository().GetByID(ctx, request.TargetUserID)
	if err != nil {
		return err
	}

	if userTarget.CreatedAt.IsZero() {
		return custerr.Forbidden(errors.New("user target no not found"))
	}

	now := time.Now().UTC()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	swipes, err := ss.repository.GetSwipeRepository().GetUserSwipe(ctx, &domain.UserSwipe{
		UserID:    user.ID,
		CreatedAt: startOfDay,
	})
	if err != nil {
		return err
	}

	if !user.CanSwipe(len(swipes)) {
		return custerr.Forbidden(errors.New("daily swipe limit reached"))
	}

	if ss.checkUserTarget(swipes, request.TargetUserID) {
		return custerr.Forbidden(errors.New("user already swiped"))
	}

	swipe := &domain.Swipe{
		ID:           ulid.Make().String(),
		UserID:       user.ID,
		TargetUserID: request.TargetUserID,
		Action:       request.Action,
		CreatedAt:    time.Now().UTC(),
	}

	return ss.repository.GetSwipeRepository().Create(ctx, swipe)
}

func (ss *SwipeService) checkUserTarget(swipes []*domain.Swipe, userTarget string) bool {
	for _, v := range swipes {
		if v.TargetUserID == userTarget {
			return true
		}
	}
	return false
}
