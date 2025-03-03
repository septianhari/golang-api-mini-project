package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/septianhari/golang-api-mini-project/models"
	"github.com/septianhari/golang-api-mini-project/repositories"
	"github.com/septianhari/golang-api-mini-project/utils"
)

func AuthRequired() func(*fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		claims, err := utils.VerifyToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		var user models.User
		repositories.GetUserById(&user, claims.UserId)

		c.Locals("user", user)

		return c.Next()
	}
}
