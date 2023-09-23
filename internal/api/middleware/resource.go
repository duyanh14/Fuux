package middleware

import "github.com/gofiber/fiber/v2"

func resource() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := _auth(c)
		if err != nil {
			return c.JSON(entity.ResponseError(err))
		}

		return c.Next()
	}
}
