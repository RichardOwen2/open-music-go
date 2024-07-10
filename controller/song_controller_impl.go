package controller

import (
	"openmusic-api/model/web"
	"openmusic-api/service"

	"github.com/gofiber/fiber/v2"
)

type SongControllerImpl struct {
	Service service.SongService
}

func NewSongController(service service.SongService) SongController {
	return &SongControllerImpl{
		Service: service,
	}
}

func (c *SongControllerImpl) Create(ctx *fiber.Ctx) error {
	var body web.SongCreateRequest

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	song, err := c.Service.Create(ctx.Context(), body)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(web.NewResponseWithData(
		fiber.StatusCreated,
		"success",
		web.SongCreateResponse{
			SongID: song.SongID,
		},
	))
}

func (c *SongControllerImpl) Update(ctx *fiber.Ctx) error {
	var body web.SongUpdateRequest
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
		"Song updated successfully",
	))
}

func (c *SongControllerImpl) Delete(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")

	if err := c.Service.Delete(ctx.Context(), id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.NewResponseWithMessage(
		fiber.StatusOK,
		"success",
		"Song deleted successfully",
	))
}

func (c *SongControllerImpl) FindById(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")

	song, err := c.Service.FindById(ctx.Context(), id)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.NewResponseWithData(
		fiber.StatusOK,
		"success",
		web.SongDataResponse{
			ID:   song.ID,
			Year: song.Year,
			Genre: song.Genre,
			Performer: song.Performer,
			Duration: song.Duration,
		},
	))
}

func (c *SongControllerImpl) FindAll(ctx *fiber.Ctx) error {
	songs, err := c.Service.FindAll(ctx.Context())

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.NewResponseWithData(
		fiber.StatusOK,
		"success",
		songs,
	))
}
