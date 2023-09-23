package api

import (
	"fuux/internal/entity"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func Run(lc fx.Lifecycle, config *entity.Config) *fiber.App {

	//lc.Append(fx.Hook{
	//	OnStart: func(ctx context.Context) error {
	//		logrus.Info(config)
	//		go api.Listen(fmt.Sprintf(":%s", config.Listen))
	//		return nil
	//	},
	//})

	return api
}
