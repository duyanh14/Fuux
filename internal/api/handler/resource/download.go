package resource

import (
	"fmt"
	"fuux/internal/api/middleware"
	"github.com/gofiber/fiber/v2"
	"os"
)

type downloadHandler struct {
}

func Download(app *fiber.App) *downloadHandler {
	handler := downloadHandler{}
	app.Get("/:resource",
		middleware.Auth,
		middleware.AllowDownload,
		handler.download)

	return &handler
}

func (h *downloadHandler) download(c *fiber.Ctx) error {

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
