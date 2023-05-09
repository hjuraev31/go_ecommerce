package middlewares

import (
	"github.com/gofiber/fiber/v2"
	helper "github.com/hjuraev31/utils"
)

func IsAuthenticated(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	if _, err := helper.ParseJWT(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "User is not authenticated!",
		})
	}
	return c.Next()
}
