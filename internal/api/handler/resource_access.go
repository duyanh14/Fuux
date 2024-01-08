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

type ResourceAccess struct {
	uc *usecase.ResourceAccess
}

func NewResourceAccess(app *fiber.App, uc *usecase.ResourceAccess) *ResourceAccess {
	handler := ResourceAccess{
		uc: uc,
	}

	app.Get("resource/access",
		middleware.Resource,
		handler.list)

	app.Get("resource/access/:id",
		middleware.Resource,
		handler.get)

	app.Post("resource/access",
		middleware.Resource,
		handler.create)

	app.Put("resource/access/:id",
		middleware.Resource,
		handler.update)

	app.Delete("resource/access/:id",
		middleware.Resource,
		handler.delete)

	return &handler
}

func (s *ResourceAccess) list(c *fiber.Ctx) error {
	var list entity.ResourceList

	err := json.Unmarshal(c.Body(), &list)
	if err != nil {
		log.Error(err)
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	rs, count, err := s.uc.List(&list)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(entity.Response{Data: fiber.Map{
		"list":  rs,
		"count": count,
	}})
}

func (s *ResourceAccess) get(c *fiber.Ctx) error {
	id := strings.Clone(c.Params("id"))

	resourceAccessModel, err := s.uc.Get(id)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.PathRecordNotFound))
	}
	return c.JSON(resourceAccessModel)
}

func (h *ResourceAccess) create(c *fiber.Ctx) error {
	payload := &entity.ResourceAccess{}

	err := c.BodyParser(payload)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	pathModel, err := h.uc.Create(payload)
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

func (h *ResourceAccess) update(c *fiber.Ctx) error {
	payload := &entity.ResourceAccess{}

	err := c.BodyParser(payload)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}
	id := strings.Clone(c.Params("id"))
	payload.ID = id

	pathModel, _, err := h.uc.UpdatePath(payload)
	if err != nil {
		exe := errorEntity.ExposeError(err,
			errorEntity.PathExist,
			errorEntity.NameExist,
			errorEntity.NameAlreadyUse,
			errorEntity.PathAlreadyUse,
		)

		return c.JSON(entity.ResponseError(exe))
	}

	return c.JSON(entity.Response{
		Error:   0,
		Message: "",
		Data:    pathModel,
	})

	//return c.JSON(entity.Response{Data: fiber.Map{
	//	"info":         account,
	//	"access_token": accessToken,
	//	//"refresh_token": refreshToken,
	//}})
}

func (h *ResourceAccess) delete(c *fiber.Ctx) error {
	id := strings.Clone(c.Params("id"))

	resourceAccessModel, err := h.uc.Get(id)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.PathRecordNotFound))
	}

	err = h.uc.RemovePath(resourceAccessModel)
	if err != nil {
		return c.JSON(entity.ResponseError(errorEntity.Unknown))
	}

	return c.JSON(entity.SuccessResponse())
}
