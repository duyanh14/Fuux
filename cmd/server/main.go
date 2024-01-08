package main

import (
	"context"
	"flag"
	apiHandler "fuux/internal/api/handler"
	"fuux/internal/entity"
	"fuux/internal/repository"
	"fuux/internal/usecase"
	"fuux/pkg"
	"fuux/pkg/database/postgres"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"log"
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

func newDatabase(config *entity.Config) *entity.Database {
	postgres, err := postgres.Connect(config.Database["postgres"])
	if err != nil {
		log.Fatalln(err)
	}

	return &entity.Database{
		Postgres: postgres,
	}
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
