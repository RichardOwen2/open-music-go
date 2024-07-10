package controller

import (
	"openmusic-api/model/web"
	"openmusic-api/service"

	"github.com/gofiber/fiber/v2"
)

type AlbumControllerImpl struct {
	Service service.AlbumService
}

func NewAlbumController(service service.AlbumService) AlbumController {
	return &AlbumControllerImpl{
		Service: service,
	}
}

func (c *AlbumControllerImpl) Create(ctx *fiber.Ctx) error {
	var body web.AlbumCreateRequest

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	album, err := c.Service.Create(ctx.Context(), body)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.NewResponseWithData(
		fiber.StatusCreated,
		"success",
		web.AlbumCreateResponse{
			AlbumID: album.AlbumID,
		},
	))
}

func (c *AlbumControllerImpl) Update(ctx *fiber.Ctx) error {
	var body web.AlbumUpdateRequest
	var id = ctx.Params("id")

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body.ID = id

	if err := c.Service.Update(ctx.Context(), body); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.NewResponseWithMessage(
		fiber.StatusOK,
		"success",
		"Album updated successfully",
	))
}

func (c *AlbumControllerImpl) Delete(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")

	if err := c.Service.Delete(ctx.Context(), id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.NewResponseWithMessage(
		fiber.StatusOK,
		"success",
		"Album deleted successfully",
	))
}

func (c *AlbumControllerImpl) FindById(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")

	album, err := c.Service.FindById(ctx.Context(), id)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.NewResponseWithData(
		fiber.StatusOK,
		"success",
		web.AlbumDataResponse{
			ID:   album.ID,
			Name: album.Name,
			Year: album.Year,
		},
	))
}

func (c *AlbumControllerImpl) FindAll(ctx *fiber.Ctx) error {
	albums, err := c.Service.FindAll(ctx.Context())

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.NewResponseWithData(
		fiber.StatusOK,
		"success",
		albums,
	))
}
