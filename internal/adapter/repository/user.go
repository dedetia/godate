package repository

import (
	"context"
	"errors"
	"github.com/dedetia/godate/internal/core/domain"
	mongostore "github.com/dedetia/golib/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const userCollection = "users"

type UserRepository struct {
	col *mongo.Collection
}

func NewUserRepository(db *mongostore.Mongo) *UserRepository {
	return &UserRepository{
		col: db.DB.Collection(userCollection),
	}
}

func (ur *UserRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*domain.User, error) {
	var user domain.User

	filter := bson.M{"phone_number": phoneNumber}

	err := ur.col.FindOne(ctx, filter).Decode(&user)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetByPhoneNumberOrEmail(ctx context.Context, phoneNumber, email string) (*domain.User, error) {
	var user domain.User

	filter := bson.M{
		"$or": []bson.M{
			{"phone_number": phoneNumber},
			{"email": email},
		},
	}

	err := ur.col.FindOne(ctx, filter).Decode(&user)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) Create(ctx context.Context, user *domain.User) error {
	_, err := ur.col.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	filter := bson.M{"_id": id}

	err := ur.col.FindOne(ctx, filter).Decode(&user)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserRecommendation(ctx context.Context, request *domain.UserRecommendation) ([]*domain.User, error) {
	filter := bson.M{
		"_id": bson.M{"$nin": request.SwipedUserIDs},
	}

	opts := options.Find().
		SetLimit(request.Limit).
		SetSkip(request.Skip)

	cur, err := ur.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, 0)
	err = cur.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) CountUserRecommendation(ctx context.Context, request *domain.UserRecommendation) (int64, error) {
	filter := bson.M{
		"_id": bson.M{"$nin": request.SwipedUserIDs},
	}
	return ur.col.CountDocuments(ctx, filter)
}

func (ur *UserRepository) Update(ctx context.Context, user *domain.User) error {
	update := bson.M{"$set": user}
	_, err := ur.col.UpdateByID(ctx, user.ID, update)
	if err != nil {
		return err
	}
	return nil
}
