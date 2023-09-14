package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	app := fiber.New()

	app.Get("/", Download)
	app.Post("/", Upload)

	app.Listen(":2210")
}

func Upload(c *fiber.Ctx) error {
	access_token := c.Query("access_token")
	if access_token != "1664661039" {
		return c.SendString("access_token")
	}
	path := c.Query("path")
	path = fmt.Sprintf("./data/%s", path)
	createDirIfNotExist(path)

	binary := c.FormValue("binary")
	if binary != "" {
		f, err := os.Create(fmt.Sprintf(path+"/%s", c.FormValue("file_name")))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err2 := f.WriteString(binary)
		if err2 != nil {
			log.Fatal(err2)
		}
		return c.SendString("ok")
	}

	form, err := c.MultipartForm()
	if err != nil {
	}
	for _, fileHeaders := range form.File {
		for _, fileHeader := range fileHeaders {
			c.SaveFile(fileHeader, fmt.Sprintf(path+"/%s", fileHeader.Filename))
		}
	}
	return c.SendString("ok")
}

func Download(c *fiber.Ctx) error {
	access_token := c.Query("access_token")
	if access_token != "1664661039" {
		return c.SendString("access_token")
	}
	path := c.Query("path")
	path = fmt.Sprintf("./data/%s", path)
	c.Attachment(path)
	return c.Download(path)
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
}
