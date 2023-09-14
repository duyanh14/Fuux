package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

type downloadHandler struct {
}

func Download(app *fiber.App) *downloadHandler {
	handler := downloadHandler{}
	app.Get("/", handler.download)

	return &handler
}

func (h *downloadHandler) download(c *fiber.Ctx) error {
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
