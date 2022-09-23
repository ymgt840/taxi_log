package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/taiki-nd/taxi_log/api/v1/controller"
)

func Routes(app *fiber.App) {

	// user
	app.Get("/api/v1/users", controller.UsersIndex)
	app.Get("/api/v1/users/:id", controller.UsersShow)
	app.Post("/api/v1/users", controller.UsersCreate)
	app.Put("/api/v1/users/:id", controller.UsersUpdate)
	app.Delete("/api/v1/users/:id", controller.UsersDelete)

	// record
	app.Get("/api/v1/records", controller.RecordsIndex)
	app.Get("/api/v1/records/:id", controller.RecordsShow)
	app.Post("/api/v1/records", controller.RecordsCreate)
	app.Put("/api/v1/records/:id", controller.RecordsUpdate)
	app.Delete("/api/v1/records/:id", controller.RecordsDelete)
}