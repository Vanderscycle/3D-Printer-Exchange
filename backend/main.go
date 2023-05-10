package main

import (
	"github.com/Vanderscycle/3D-Printer-Exchange/database"
	"github.com/Vanderscycle/3D-Printer-Exchange/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"

	"github.com/gofiber/swagger"
	// docs are generated by Swag CLI, you have to import them.
	// replace with your own docs folder, usually "github.com/username/reponame/docs"
	_ "github.com/Vanderscycle/3D-Printer-Exchange/docs"
)

//	@title			Printer Exchange Api
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

// @contact.name	Henri Vandersleyen
// @contact.url	http://www.swagger.io/support
// @contact.email	henri-vandersleyen@protonmail.com
// @license.name	BSD-3-Clause license
// @license.url	https://opensource.org/license/bsd-3-clause/
func main() {

	database.Connect()

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Use(logger.New())

	app.Use(cors.New())

	//Routes layout
	router.SetupUserRoutes(app)
	router.SetupPrinterRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":8080")
}
