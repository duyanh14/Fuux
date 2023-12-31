package resource

import (
	"fmt"
	"fuux/internal/api/middleware"
	"github.com/gofiber/fiber/v2"
	"os"
)

type uploadHandler struct {
}

func Upload(app *fiber.App) *uploadHandler {
	handler := uploadHandler{}
	app.Post("/:resource",
		middleware.Auth,
		middleware.AllowUpload,
		handler.upload)

	return &handler
}

func (h *uploadHandler) upload(c *fiber.Ctx) error {

	path := c.Query("path")
	path = fmt.Sprintf("./data/%s", path)
	createDirIfNotExist(path)

	binary := c.FormValue("binary")
	if binary != "" {
		f, err := os.Create(fmt.Sprintf(path+"/%s", c.FormValue("file_name")))
		if err != nil {
			return err
		}
		defer f.Close()
		_, err2 := f.WriteString(binary)
		if err2 != nil {
			return err
		}
		return c.SendString("ok")
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	for _, fileHeaders := range form.File {
		for _, fileHeader := range fileHeaders {
			c.SaveFile(fileHeader, fmt.Sprintf(path+"/%s", fileHeader.Filename))
		}
	}
	return c.SendString("ok")
}
