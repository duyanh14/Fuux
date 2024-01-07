package main

import (
	"context"
	"flag"
	"fuux/internal/api/handler"
	"fuux/internal/entity"
	"fuux/internal/repository"
	resourceRepository "fuux/internal/repository/resource"
	"fuux/internal/usecase"
	"fuux/internal/usecase/resource"
	"fuux/pkg"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var (
	flagConf string
)

func init() {
	flag.StringVar(&flagConf, "conf", "prod", "config path, eg: -conf config.yaml")
}

func newConfig() (*entity.Config, error) {
	return pkg.NewConfig(flagConf)
}

func newApp(lc fx.Lifecycle) *fiber.App {
	app := fiber.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(":3000")
			return nil
		},
	})

	return app
}

func main() {
	flag.Parse()

	fx.New(
		fx.Provide(
			newConfig,
			usecase.NewDatabase,
			resourceRepository.NewResource,
			resourceRepository.NewResourceAccess,
			repository.NewFile,
			handler.Resource,
			handler.ResourceAccess,
			handler.Download,
			handler.Upload,
			resource.NewResource,
			resource.NewResourceAccess,
			newApp),
		fx.Invoke(func(*fiber.App) {}),
	).Run()
}
