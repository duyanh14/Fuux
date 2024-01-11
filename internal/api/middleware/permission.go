package middleware

import (
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	"fuux/pkg"
	"github.com/gofiber/fiber/v2"
)

func internalPermission(config *entity.Config) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ip := c.Get("X-Real-IP")

		intIP, err := pkg.InternalIP(ip)
		if err != nil {
			return c.JSON(entity.ResponseError(errorEntity.PermissionDenied))
		}

		if intIP {
			return c.Next()
		}

		return c.JSON(entity.ResponseError(errorEntity.PermissionDenied))
	}
}

func allowUpload() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := _allowUpload(c)
		if err != nil {
			return c.JSON(entity.ResponseError(err))
		}

		return c.Next()
	}
}

func _allowUpload(c *fiber.Ctx) *errorEntity.Error {
	model := c.Locals("resource_access").(*entity.ResourceAccess)

	if model.Status == entity.ResourceAccessStatusDisable {
		return errorEntity.ResourceAccessStatusIsDisable
	}

	if model.Permission.Write == false {
		return errorEntity.UploadDisallow
	}

	return nil
}
func allowDownload() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := _allowDownload(c)
		if err != nil {
			return c.JSON(entity.ResponseError(err))
		}

		return c.Next()
	}
}

func _allowDownload(c *fiber.Ctx) *errorEntity.Error {
	model := c.Locals("resource_access").(*entity.ResourceAccess)

	if model.Status == entity.ResourceAccessStatusDisable {
		return errorEntity.ResourceAccessStatusIsDisable
	}

	if model.Permission.Read == false {
		return errorEntity.DownloadDisallow
	}

	return nil
}
