package main

import (
	"flag"
	"fuux/internal/api/handler/resource"
	"fuux/internal/entity"
	"fuux/internal/repository"
	service "fuux/internal/usecase"
	"fuux/pkg"
	"github.com/gofiber/fiber/v2"
	"log"
)

func f() *entity.Flag {
	f := &entity.Flag{
		Config: flag.String("conf", "", "config resource"),
	}
	flag.Parse()

	if f.Config == nil {
		log.Fatal("no config resource")
	}

	return f
}

func main() {
	app := fiber.New()

	env := flag.String("env", "dev", "Environment")

	flag.Parse()

	if *env == "" {
		log.Fatal("No environment")
	}

	config := pkg.NewConfig(*env)

	db, err := service.NewDatabase(config)
	if err != nil {
		log.Fatal(err)
	}

	repository.File, err = repository.NewFile(db)
	if err != nil {
		return
	}
	resource.Path(app)
	resource.Download(app)
	resource.Upload(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
