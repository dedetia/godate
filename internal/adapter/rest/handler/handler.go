package handler

import (
	"github.com/dedetia/godate/internal/core/port/registry"
)

type Handler struct {
	service registry.ServiceRegistry
}

func NewHandler(service registry.ServiceRegistry) *Handler {
	return &Handler{
		service: service,
	}
}
