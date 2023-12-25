package main

import (
	"flag"
	"fuux/internal/api/handler/resource"
	"fuux/internal/entity"
	"fuux/internal/repository"
	resourceRepository "fuux/internal/repository/resource"
	"fuux/internal/usecase"
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

	initDB := entity.Database{
		Postgres:  db,
		SQLServer: nil,
	}

	resourceRepository.Resource, err = resourceRepository.NewResource(&initDB)
	if err != nil {
		return
	}

	resourceRepository.ResourceAccess, err = resourceRepository.NewResourceAccess(&initDB)
	if err != nil {
		return
	}

	repository.File, err = repository.NewFile(db)
	if err != nil {
		return
	}

	resource.Path(app)
	resource.Download(app)
	resource.Upload(app)

	usecase.NewResource(config)
	usecase.NewResourceAccess(config)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
