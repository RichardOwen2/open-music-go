package main

import (
	"openmusic-api/app"
	"openmusic-api/controller"
	"openmusic-api/exception"
	"openmusic-api/helper"
	"openmusic-api/repository"
	"openmusic-api/routes"
	"openmusic-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fiber := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	db := app.OpenConnection()
	validate := validator.New()

	songRepository := repository.NewSongRepositoryImpl()
	songService := service.NewSongServiceImpl(songRepository, db, validate)
	songController := controller.NewSongController(songService)

	albumRepository := repository.NewAlbumRepositoryImpl()
	albumService := service.NewAlbumServiceImpl(albumRepository, db, validate)
	albumController := controller.NewAlbumController(albumService)

	routes.SongRoutes(fiber, songController)
	routes.AlbumRoutes(fiber, albumController)

	err := fiber.Listen(":3000")
	helper.PanicIfError(err)
}
