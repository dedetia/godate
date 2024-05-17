package handler

import (
	"github.com/dedetia/godate/internal/core/domain"
	"github.com/dedetia/godate/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) PurchasePremium(c *fiber.Ctx) error {
	var request domain.PurchasePremium

	if err := c.BodyParser(&request); err != nil {
		return response.Error(c, fiber.ErrBadRequest, err.Error())
	}

	err := h.service.GetPurchaseService().PurchasePremium(c.Context(), &request)
	if err != nil {
		return response.AuthError(c, err)
	}

	return response.Success(c, fiber.StatusOK)
}
