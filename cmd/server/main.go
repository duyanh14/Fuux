package main

import (
	"flag"
	"fuux/internal/api"
	"fuux/internal/api/handler"
	"fuux/internal/entity"
	service "fuux/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/fx"
)

func f() *entity.Flag {
	f := &entity.Flag{
		Config: flag.String("conf", "", "Config file"),
	}
	flag.Parse()

	if f.Config == nil {
		log.Fatal("no config file")
	}
	return f
}

func main() {
	fx.New(
		fx.Provide(
			f,
			service.NewConfig,
			fiber.New,
			handler.Download,
			handler.Upload,
			api.Run),
		fx.Invoke(func(*fiber.App) {}),
	).Run()
}
