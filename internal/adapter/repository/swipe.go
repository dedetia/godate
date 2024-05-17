package repository

import (
	"context"
	"github.com/dedetia/godate/internal/core/domain"
	mongostore "github.com/dedetia/golib/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const swipeCollection = "swipes"

type SwipeRepository struct {
	col *mongo.Collection
}

func NewSwipeRepository(db *mongostore.Mongo) *SwipeRepository {
	return &SwipeRepository{
		col: db.DB.Collection(swipeCollection),
	}
}

func (ur *SwipeRepository) GetUserSwipe(ctx context.Context, request *domain.UserSwipe) ([]*domain.Swipe, error) {
	var swipes []*domain.Swipe

	filter := bson.M{
		"user_id": request.UserID,
		"created_at": bson.M{
			"$gte": request.CreatedAt,
		},
	}

	cur, err := ur.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &swipes)
	if err != nil {
		return nil, err
	}

	return swipes, nil
}

func (ur *SwipeRepository) Create(ctx context.Context, request *domain.Swipe) error {
	_, err := ur.col.InsertOne(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
