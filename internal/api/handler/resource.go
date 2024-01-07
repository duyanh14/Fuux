package handler

import (
	"encoding/json"
	"fuux/internal/api/middleware"
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	resourceRepository "fuux/internal/repository/resource"
	"fuux/internal/usecase/resource"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"strings"
)

type resourceHandler struct {
}

func Resource(app *fiber.App) *resourceHandler {
	handler := resourceHandler{}
	app.Get("resource/:id",
		middleware.Resource,
		handler.get)
	app.Get("resource",
		middleware.Resource,
		handler.getList)

	app.Post("resource",
		middleware.Resource,
		handler.create)

	app.Put("resource/:id",
		middleware.Resource,
		handler.update)

	app.Delete("resource/:id",
		middleware.Resource,
		handler.delete)

	return &handler
}

func (h *resourceHandler) update(c *fiber.Ctx) error {
	payload := &entity.Resource{}

	err := c.BodyParser(payload)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}
	id := strings.Clone(c.Params("id"))
	payload.ID = id

	model, _, err := resource.Resource.UpdatePath(payload)
	if err != nil {
		exe := errorEntity.ExposeError(err,
			errorEntity.PathExist,
			errorEntity.NameExist,
			errorEntity.NameAlreadyUse,
			errorEntity.PathAlreadyUse,
		)

		return c.JSON(entity.ResponseError(exe))
	}

	return c.JSON(model)

	//return c.JSON(entity.Response{Data: fiber.Map{
	//	"info":         account,
	//	"access_token": accessToken,
	//	//"refresh_token": refreshToken,
	//}})
}

func (h *resourceHandler) create(c *fiber.Ctx) error {
	payload := &entity.Resource{}

	err := c.BodyParser(payload)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	model, _, err := resource.Resource.AddResource(payload)
	if err != nil {
		exe := errorEntity.ExposeError(err,
			errorEntity.FieldRequired,
			errorEntity.NameAlreadyUse,
			errorEntity.PathAlreadyUse,
		)

		return c.JSON(entity.ResponseError(exe))
	}

	return c.JSON(model)

	//return c.JSON(entity.Response{Data: fiber.Map{
	//	"info":         account,
	//	"access_token": accessToken,
	//	//"refresh_token": refreshToken,
	//}})
}

func (s *resourceHandler) delete(c *fiber.Ctx) error {
	id := strings.Clone(c.Params("id"))

	model, err := resourceRepository.Resource.GetByID(id)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.PathRecordNotFound))
	}

	err = resource.Resource.RemovePath(model)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	return c.JSON(entity.SuccessResponse())
}

func (s *resourceHandler) getList(c *fiber.Ctx) error {
	var list entity.ResourceList

	err := json.Unmarshal(c.Body(), &list)
	if err != nil {
		log.Error(err)
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	rs, count, err := resource.Resource.List(&list)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(entity.Response{Data: fiber.Map{
		"list":  rs,
		"count": count,
	}})
}

func (s *resourceHandler) get(c *fiber.Ctx) error {
	id := strings.Clone(c.Params("id"))

	model, err := resource.Resource.Get(id)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.PathRecordNotFound))
	}
	return c.JSON(model)
}
