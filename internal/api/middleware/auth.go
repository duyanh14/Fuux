package middleware

import (
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	"fuux/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func getToken(c *fiber.Ctx) (string, *errorEntity.Error) {
	token := strings.Replace(c.GetReqHeaders()["Authorization"], "Bearer ", "", -1)
	if token == "" {
		return "", errorEntity.PermissionDenied
	}
	return token, nil
}

func auth() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := _auth(c)
		if err != nil {
			return c.JSON(entity.ResponseError(err))
		}

		return c.Next()
	}
}

func _auth(c *fiber.Ctx) *errorEntity.Error {
	accessToken, er := getToken(c)
	if er != nil {
		return er
	}

	model, err := usecase.Upload.Auth(accessToken)
	if err != nil {
		return errorEntity.PermissionDenied
	}

	c.Locals("resource_access", model)

	return nil
}
