package router

import (
	"github.com/Vanderscycle/3D-Printer-Exchange/handler"
	"github.com/gofiber/fiber/v2"
) // SetupRoutes func
func SetupUserRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/user") // routes
	v1.Get("/", handler.GetAllUsers)
	v1.Get("/:id", handler.GetSingleUser)
	v1.Post("/", handler.CreateUser)
	v1.Put("/:id", handler.UpdateUser)
	v1.Delete("/:id", handler.DeleteUserByID)
}

func SetupPrinterRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/printer") // routes
	v1.Get("/", handler.GetAllPrinters)
	v1.Get("/:id", handler.GetSinglePrinter)
	v1.Post("/", handler.CreatePrinter)
	v1.Put("/:id", handler.UpdatePrinter)
	v1.Delete("/:id", handler.DeletePrinterByID)
}
