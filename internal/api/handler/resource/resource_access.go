package resource

import (
	"encoding/json"
	"fuux/internal/api/middleware"
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	"fuux/internal/repository"
	resourceRepository "fuux/internal/repository/resource"
	"fuux/internal/usecase"
	"fuux/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"path/filepath"
	"strings"
)

type pathHandler struct {
}

func Path(app *fiber.App) *pathHandler {
	handler := pathHandler{}
	app.Get("path/:id",
		middleware.Resource,
		handler.get)
	app.Get("path",
		middleware.Resource,
		handler.getList)

	app.Post("path",
		middleware.Resource,
		handler.addPath)

	app.Put("path/:id",
		middleware.Resource,
		handler.updatePath)

	app.Delete("path/:id",
		middleware.Resource,
		handler.removePath)

	return &handler
}

func (h *pathHandler) updatePath(c *fiber.Ctx) error {
	payload := &entity.Path{}

	err := c.BodyParser(payload)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}
	id := strings.Clone(c.Params("id"))
	payload.ID = id

	pathModel, _, err := usecase.Resource.UpdatePath(payload)
	if err != nil {
		exe := errorEntity.ExposeError(err,
			errorEntity.PathExist,
			errorEntity.NameExist,
			errorEntity.NameAlreadyUse,
			errorEntity.PathAlreadyUse,
		)

		return c.JSON(entity.ResponseError(exe))
	}

	return c.JSON(pathModel)

	//return c.JSON(entity.Response{Data: fiber.Map{
	//	"info":         account,
	//	"access_token": accessToken,
	//	//"refresh_token": refreshToken,
	//}})
}

func (h *pathHandler) addPath(c *fiber.Ctx) error {
	payload := &entity.Path{}

	err := c.BodyParser(payload)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	pathModel, _, err := usecase.Resource.AddPath(payload)
	if err != nil {
		exe := errorEntity.ExposeError(err,
			errorEntity.FieldRequired,
			errorEntity.NameAlreadyUse,
			errorEntity.PathAlreadyUse,
		)

		return c.JSON(entity.ResponseError(exe))
	}

	return c.JSON(pathModel)

	//return c.JSON(entity.Response{Data: fiber.Map{
	//	"info":         account,
	//	"access_token": accessToken,
	//	//"refresh_token": refreshToken,
	//}})
}

func (s *pathHandler) removePath(c *fiber.Ctx) error {
	id := strings.Clone(c.Params("id"))

	pathModel, err := resourceRepository.Resource.GetByID(id)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.PathRecordNotFound))
	}

	err = usecase.Resource.RemovePath(pathModel)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	return c.JSON(entity.SuccessResponse())
}

func (s *pathHandler) getList(c *fiber.Ctx) error {
	var list entity.PathList

	err := json.Unmarshal(c.Body(), &list)
	if err != nil {
		log.Error(err)
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	rs, count, err := usecase.Resource.List(&list)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(entity.Response{Data: fiber.Map{
		"list":  rs,
		"count": count,
	}})
}

func (s *pathHandler) get(c *fiber.Ctx) error {
	id := strings.Clone(c.Params("id"))

	pathModel, err := usecase.Resource.Get(id)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.PathRecordNotFound))
	}
	return c.JSON(pathModel)
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

	name := data["name"]
	path := filepath.Dir(data["path"])

	var pathModel = entity.Path{
		ID:   uuid.NewString(),
		Name: name,
		Path: path,
	}

	pathExist, _ := repository.MatchRecord("path", path, &entity.Path{})
	if pathExist == true {
		return c.JSON(fiber.Map{"error": "path already exists"})
	}

	result := repository.File.Db.Model(&entity.Path{}).Create(&pathModel)

	if result.Error != nil {
		return c.JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(pathModel)

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
		return c.JSON(fiber.Map{"error": "can not found your path"})
	}

	pathModel := rawPath.(*entity.Path)

	if pathModel.ID != data["id"].(string) && filepath.Dir(pathModel.Path) == filepath.Dir(data["path"].(string)) {
		return c.JSON(fiber.Map{"error": "this path exist in database"})
	}

	rs := repository.File.Db.Model(&pathModel).Updates(data)
	if rs.Error != nil {
		return c.JSON(fiber.Map{"error": rs.Error.Error()})
	}
	if rs.RowsAffected == 1 {
		return c.JSON(fiber.Map{"success": "true"})
	}

	return nil
}
func (h *pathHandler) getOld(c *fiber.Ctx) error {

	id := c.Params("resource")

	pathExist, rawPath := repository.MatchRecord("id", id, &entity.Path{})
	if pathExist == false {
		return c.JSON(fiber.Map{"error": "path does not exists"})
	}
	if rawPath == nil {
		return c.JSON(fiber.Map{"error": "can not found your path"})
	}

	path := rawPath.(*entity.Path)

	return c.JSON(path)

}
func (h *pathHandler) getListOld(c *fiber.Ctx) error {
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
	pathName := c.Params("resource")

	pathExist, rawPath := repository.MatchRecord("path", pathName, &entity.Path{})
	if pathExist == false {
		return c.JSON(fiber.Map{"error": "path does not exists"})
	}
	if rawPath == nil {
		return c.JSON(fiber.Map{"error": "can not found your path"})
	}

	path := pathExist.(*entity.Path)

	var rs = repository.File.Db.Delete(&path)

	if rs.Error != nil {
		return c.JSON(fiber.Map{"error": rs.Error.Error()})
	}
	if rs.RowsAffected == 1 {
		return c.JSON(fiber.Map{"success": "true"})
	}
	return nil
}
