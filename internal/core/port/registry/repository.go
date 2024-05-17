package registry

import (
	"context"
	"github.com/dedetia/godate/internal/core/port/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryRegistry interface {
	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) (interface{}, error)) error
	GetUserRepository() repository.UserRepository
	GetSwipeRepository() repository.SwipeRepository
}
