package middleware

import (
	"github.com/dedetia/godate/pkg/auth"
	"github.com/dedetia/godate/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return response.Error(c, fiber.ErrUnauthorized)
	}

	claim, err := auth.ClaimJWT(token)
	if err != nil {
		return response.Error(c, fiber.ErrUnauthorized, err.Error())
	}

	c.Locals(auth.UserKey, claim)

	return c.Next()
}
