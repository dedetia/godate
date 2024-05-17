package domain

import "time"

type (
	UserSwipe struct {
		UserID    string
		CreatedAt time.Time
	}

	Swipe struct {
		ID           string    `bson:"_id"`
		UserID       string    `bson:"user_id"`
		TargetUserID string    `bson:"target_user_id"`
		Action       string    `bson:"action"`
		CreatedAt    time.Time `bson:"created_at"`
	}

	SwipeRequest struct {
		TargetUserID string `json:"target_user_id" validate:"required"`
		Action       string `json:"action" validate:"required,oneof=like pass"`
	}
)
