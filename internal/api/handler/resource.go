package handler

import (
	"encoding/json"
	"fuux/internal/api/middleware"
	"fuux/internal/entity"
	errorEntity "fuux/internal/entity/error"
	"fuux/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"strings"
)

type Resource struct {
	uc    *usecase.Resource
	group fiber.Router
}

func NewResource(app *fiber.App, uc *usecase.Resource) *Resource {
	group := app.Group("/resource", middleware.InternalPermission)

	handler := Resource{
		uc:    uc,
		group: group,
	}

	group.Get("/",
		middleware.Resource,
		handler.list)

	group.Get("/:id",
		middleware.Resource,
		handler.get)

	group.Post("/",
		middleware.Resource,
		handler.create)

	group.Put("/:id",
		middleware.Resource,
		handler.update)

	group.Delete("/:id",
		middleware.Resource,
		handler.delete)

	return &handler
}

func (h *Resource) list(c *fiber.Ctx) error {
	var list entity.ResourceList

	err := json.Unmarshal(c.Body(), &list)
	if err != nil {
		log.Error(err)
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	rs, count, err := h.uc.List(&list)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(entity.Response{Data: fiber.Map{
		"list":  rs,
		"count": count,
	}})
}

func (h *Resource) get(c *fiber.Ctx) error {

	id := strings.Clone(c.Params("id"))

	model, err := h.uc.Get(id)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.PathRecordNotFound))
	}
	return c.JSON(model)
}

func (h *Resource) create(c *fiber.Ctx) error {
	payload := &entity.Resource{}

	err := c.BodyParser(payload)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	model, _, err := h.uc.AddResource(payload)
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

func (h *Resource) update(c *fiber.Ctx) error {
	payload := &entity.Resource{}

	err := c.BodyParser(payload)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}
	id := strings.Clone(c.Params("id"))
	payload.ID = id

	model, _, err := h.uc.UpdatePath(payload)
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

func (h *Resource) delete(c *fiber.Ctx) error {
	id := strings.Clone(c.Params("id"))

	model, err := h.uc.Get(id)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.PathRecordNotFound))
	}

	err = h.uc.RemovePath(model)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	return c.JSON(entity.SuccessResponse())
}
