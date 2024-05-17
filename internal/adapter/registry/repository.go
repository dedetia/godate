package registry

import (
	"context"
	"github.com/dedetia/godate/internal/adapter/repository"
	"github.com/dedetia/godate/internal/core/port/registry"
	portrepository "github.com/dedetia/godate/internal/core/port/repository"
	mongostore "github.com/dedetia/golib/storage/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryRegistry struct {
	db              *mongostore.Mongo
	userRepository  portrepository.UserRepository
	swipeRepository portrepository.SwipeRepository
}

func NewRepositoryRegistry(db *mongostore.Mongo) registry.RepositoryRegistry {
	return &RepositoryRegistry{
		db:              db,
		userRepository:  repository.NewUserRepository(db),
		swipeRepository: repository.NewSwipeRepository(db),
	}
}

func (rr *RepositoryRegistry) WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) (interface{}, error)) error {
	ses, err := rr.db.DB.Client().StartSession()
	if err != nil {
		return err
	}
	defer ses.EndSession(ctx)

	_, err = ses.WithTransaction(ctx, fn)
	if err != nil {
		return err
	}

	return nil
}

func (rr *RepositoryRegistry) GetUserRepository() portrepository.UserRepository {
	return rr.userRepository
}

func (rr *RepositoryRegistry) GetSwipeRepository() portrepository.SwipeRepository {
	return rr.swipeRepository
}
