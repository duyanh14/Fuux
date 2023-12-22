package resource

import (
	"fmt"
	"fuux/internal/api/middleware"
	"fuux/internal/entity"
	"fuux/internal/repository"
	"fuux/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type pathHandler struct {
}

func Path(app *fiber.App) *pathHandler {
	handler := pathHandler{}
	app.Get("path/:resource",
		middleware.Resource,
		handler.get)
	app.Get("path/",
		middleware.Resource,
		handler.getList)

	app.Post("path/",
		middleware.Resource,
		handler.create)

	app.Put("path/",
		middleware.Resource,
		handler.update)

	app.Delete("path/",
		middleware.Resource,
		handler.delete)

	return &handler
}

func (h *pathHandler) create(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	checkParams := pkg.AllKeyRequired(data, []string{"name", "path"})

	if checkParams == false {
		return c.JSON(fiber.Map{"error": "missing required parameter"})
	}

	var path = entity.Path{
		ID:   uuid.NewString(),
		Name: data["name"],
		Path: data["path"],
	}

	pathExist, _ := repository.MatchRecord("path", data["path"], &entity.Path{})
	if pathExist == true {
		return c.JSON(fiber.Map{"error": "path already exists"})
	}

	result := repository.File.Db.Model(&entity.Path{}).Create(&path)

	if result.Error != nil {
		return c.JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(path)

}
func (h *pathHandler) update(c *fiber.Ctx) error {
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	nilValidate := pkg.IsMapContainNil(data, []string{"name", "path"})

	if nilValidate {
		return c.JSON(fiber.Map{"error": "you send empty value"})
	}

	pathExist, rawPath := repository.MatchRecord("id", data["id"], &entity.Path{})

	if pathExist == false {
		return c.JSON(fiber.Map{"error": "path does not exists"})
	}
	if rawPath == nil {
		return c.JSON(fiber.Map{"error": "can not found your user"})
	}

	path := rawPath.(*entity.Path)

	rs := repository.File.Db.Model(&path).Updates(data)
	if rs.Error != nil {
		return c.JSON(fiber.Map{"error": rs.Error.Error()})
	}
	if rs.RowsAffected == 1 {
		return c.JSON(fiber.Map{"success": "true"})
	}

	return nil
}
func (h *pathHandler) get(c *fiber.Ctx) error {
	access_token := c.Query("access_token")
	if access_token != "1664661039" {
		return c.SendString("access_token")
	}
	path := c.Query("path")
	path = fmt.Sprintf("./data/%s", path)
	c.Attachment(path)
	return c.Download(path)
}
func (h *pathHandler) getList(c *fiber.Ctx) error {
	var records = []entity.Path{}

	rs := repository.File.Db.Find(&records)
	if rs.Error != nil {
		return c.JSON(fiber.Map{"error": rs.Error.Error()})
	}
	if rs.RowsAffected > 0 {
		return c.JSON(records)
	}
	return nil
}
func (h *pathHandler) delete(c *fiber.Ctx) error {
	access_token := c.Query("access_token")
	if access_token != "1664661039" {
		return c.SendString("access_token")
	}
	path := c.Query("path")
	path = fmt.Sprintf("./data/%s", path)
	c.Attachment(path)
	return c.Download(path)
}
