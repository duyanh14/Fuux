package main

import (
	"flag"
	"fuux/internal/api/handler/resource"
	"fuux/internal/entity"
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

	//f := f()
	//
	//config, err := pkg.Config(f)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//db, err := service.NewDatabase(config)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//repository.NewFile(db)

	resource.Download(app)
	resource.Upload(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
