package registry

import "github.com/dedetia/godate/internal/core/port/service"

type ServiceRegistry interface {
	GetUserService() service.UserService
	GetSwipeService() service.SwipeService
	GetPurchaseService() service.PurchaseService
}
