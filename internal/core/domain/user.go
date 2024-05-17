package domain

import (
	"github.com/dedetia/godate/shared/constant"
	"time"
)

type (
	LoginRequest struct {
		PhoneNumber string `json:"phone_number" validate:"required"`
		Password    string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		ID          string `json:"id"`
		AccessToken string `json:"access_token"`
	}

	SignupRequest struct {
		PhoneNumber string  `form:"phone_number" validate:"required,min=12,max=15,phone_number"`
		Password    string  `form:"password" validate:"required,min=6,max=64,uppercase,digit,special"`
		Email       string  `form:"email" validate:"required,email"`
		Name        string  `form:"name" validate:"required,min=3,max=30"`
		DOB         string  `form:"dob" validate:"required,dob"`
		Gender      string  `form:"gender" validate:"required,oneof=Male Female"`
		Photos      []*File `form:"photos" validate:"required"`
	}

	SignupResponse struct {
		ID string `json:"id"`
	}

	User struct {
		ID          string    `bson:"_id" json:"id"`
		PhoneNumber string    `bson:"phone_number" json:"phone_number"`
		Password    []byte    `bson:"password" json:"-"`
		Email       string    `bson:"email" json:"email"`
		Name        string    `bson:"name" json:"name"`
		DOB         string    `bson:"dob" json:"dob"`
		Age         int       `bson:"age" json:"age"`
		Gender      string    `bson:"gender" json:"gender"`
		Photos      []*File   `bson:"photos" json:"photos"`
		Feature     Feature   `bson:"feature,omitempty" json:"feature"`
		CreatedAt   time.Time `bson:"created_at" json:"created_at"`
		UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
	}

	Feature struct {
		IsPremium      bool   `bson:"is_premium" json:"is_premium"`
		IsVerified     bool   `bson:"is_verified" json:"is_verified"`
		PremiumFeature string `bson:"premium_feature,omitempty" json:"premium_feature"`
	}

	File struct {
		Name string `bson:"name,omitempty" json:"name"`
		Path string `bson:"path,omitempty" json:"path"`
	}

	ProfileRequest struct {
		Page int `query:"page"`
	}

	Profile struct {
		ID          string   `json:"id"`
		PhoneNumber string   `json:"phone_number"`
		Email       string   `json:"email"`
		Name        string   `json:"name"`
		DOB         string   `json:"dob"`
		Age         int      `json:"age"`
		Gender      string   `json:"gender"`
		Photos      []string `json:"photos"`
		Feature     Feature  `json:"feature"`
	}

	UserRecommendation struct {
		SwipedUserIDs []string
		Limit         int64
		Skip          int64
	}

	ProfileResponse struct {
		Profiles   []*Profile `json:"profiles"`
		Pagination Pagination `json:"pagination"`
	}
)

func (u *User) CanSwipe(count int) bool {
	if u.Feature.IsPremium {
		if u.Feature.PremiumFeature == constant.PremiumPackageVerifiedLabel.String() && count == 10 {
			return false
		}

		if u.Feature.PremiumFeature == constant.PremiumPackageNoSwipeQuota.String() {
			return true
		}
	}

	if count == 10 {
		return false
	}

	return true
}
