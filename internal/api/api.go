package api

import (
	"context"
	"fmt"
	"fuux/internal/entity"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func Run(lc fx.Lifecycle, config *entity.Config) *fiber.App {
	api := fiber.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go api.Listen(fmt.Sprintf(":%s", config.Listen))
			return nil
		},
	})

	return api
}
