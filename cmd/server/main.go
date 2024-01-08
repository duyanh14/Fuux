package main

import (
	"context"
	"flag"
	"fmt"
	apiHandler "fuux/internal/api/handler"
	"fuux/internal/entity"
	"fuux/internal/repository"
	"fuux/internal/usecase"
	"fuux/pkg"
	"fuux/pkg/database/postgres"
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

func newApp(lc fx.Lifecycle, config *entity.Config) *fiber.App {
	app := fiber.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(fmt.Sprintf(":%d", config.Listen))
			return nil
		},
	})

	return app
}

func newDatabase(config *entity.Config) (*entity.Database, error) {
	postgres, err := postgres.Connect(config.Database["postgres"])
	if err != nil {
		return nil, err
	}

	return &entity.Database{
		Postgres: postgres,
	}, nil
}

func main() {
	flag.Parse()

	fx.New(
		fx.Provide(
			newConfig,
			newDatabase,
			repository.NewFile,
			repository.NewResource,
			repository.NewResourceAccess,
			usecase.NewFile,
			usecase.NewResource,
			usecase.NewResourceAccess,
			apiHandler.NewResource,
			apiHandler.NewResourceAccess,
			apiHandler.NewFile,
			newApp),
		fx.Invoke(func(*fiber.App) {}),
	).Run()
}
