package registry

import (
	"github.com/dedetia/godate/config"
	"github.com/dedetia/godate/internal/core/port/registry"
	portservice "github.com/dedetia/godate/internal/core/port/service"
	"github.com/dedetia/godate/internal/core/service"
)

type ServiceRegistry struct {
	userService     portservice.UserService
	swipeService    portservice.SwipeService
	purchaseService portservice.PurchaseService
}

func NewServiceRegistry(cfg *config.MainConfig, repository registry.RepositoryRegistry) registry.ServiceRegistry {
	return &ServiceRegistry{
		userService:     service.NewUserService(cfg, repository),
		swipeService:    service.NewSwipeService(repository),
		purchaseService: service.NewPurchaseService(repository),
	}
}

func (sr *ServiceRegistry) GetUserService() portservice.UserService {
	return sr.userService
}

func (sr *ServiceRegistry) GetSwipeService() portservice.SwipeService {
	return sr.swipeService
}

func (sr *ServiceRegistry) GetPurchaseService() portservice.PurchaseService {
	return sr.purchaseService
}
