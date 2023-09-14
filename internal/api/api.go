package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func Run(lc fx.Lifecycle, api *fiber.App) *fiber.App {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return api.Listen(":2210")
		},
	})

	return api
}
