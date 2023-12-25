package resource

import (
	"encoding/json"
	"fuux/internal/api/middleware"
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	resourceRepository "fuux/internal/repository/resource"
	"fuux/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"strings"
)

type resourceAccessHandler struct {
}

func ResourceAccess(app *fiber.App) *resourceAccessHandler {
	handler := resourceAccessHandler{}
	app.Get("path/access/:id",
		middleware.Resource,
		handler.get)
	app.Get("path/access",
		middleware.Resource,
		handler.getList)

	app.Post("path/access",
		middleware.Resource,
		handler.addPath)

	app.Put("path/access/:id",
		middleware.Resource,
		handler.updatePath)

	app.Delete("path/access/:id",
		middleware.Resource,
		handler.removePath)

	return &handler
}

func (h *resourceAccessHandler) updatePath(c *fiber.Ctx) error {
	payload := &entity.ResourceAccess{}

	err := c.BodyParser(payload)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}
	id := strings.Clone(c.Params("id"))
	payload.ID = id

	pathModel, _, err := usecase.ResourceAccess.UpdatePath(payload)
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

func (h *resourceAccessHandler) addPath(c *fiber.Ctx) error {
	payload := &entity.ResourceAccess{}

	err := c.BodyParser(payload)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	pathModel, err := usecase.ResourceAccess.Create(payload)
	if err != nil {
		exe := errorEntity.ExposeError(err,
			errorEntity.FieldRequired,
			errorEntity.NameAlreadyUse,
			errorEntity.PathAlreadyUse,
		)

		return c.JSON(entity.ResponseError(exe))
	}

	return c.JSON(entity.Response{
		Data: pathModel,
	})

	//return c.JSON(entity.Response{Data: fiber.Map{
	//	"info":         account,
	//	"access_token": accessToken,
	//	//"refresh_token": refreshToken,
	//}})
}

func (s *resourceAccessHandler) removePath(c *fiber.Ctx) error {
	id := strings.Clone(c.Params("id"))

	resourceAccessModel, err := resourceRepository.ResourceAccess.GetByID(id)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.PathRecordNotFound))
	}

	err = usecase.ResourceAccess.RemovePath(resourceAccessModel)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	return c.JSON(entity.SuccessResponse())
}

func (s *resourceAccessHandler) getList(c *fiber.Ctx) error {
	var list entity.ResourceList

	err := json.Unmarshal(c.Body(), &list)
	if err != nil {
		log.Error(err)
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	rs, count, err := usecase.ResourceAccess.List(&list)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(entity.Response{Data: fiber.Map{
		"list":  rs,
		"count": count,
	}})
}

func (s *resourceAccessHandler) get(c *fiber.Ctx) error {
	id := strings.Clone(c.Params("id"))

	resourceAccessModel, err := usecase.ResourceAccess.Get(id)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.PathRecordNotFound))
	}
	return c.JSON(resourceAccessModel)
}
