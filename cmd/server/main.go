package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) {
		file, err := c.FormFile("document")
		if err == nil {
			c.SaveFile(file, fmt.Sprintf("./data/%s", file.Filename))
			c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
			c.SaveFile(file, fmt.Sprintf("/tmp/uploads_relative/%s", file.Filename))
		}
	})

	log.Fatal(app.Listen(3000))
}