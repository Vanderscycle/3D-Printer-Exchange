package handler

import (
	_ "github.com/Vanderscycle/3D-Printer-Exchange/response"
	"github.com/gofiber/fiber/v2"
)

// GetVersionApi	godoc
//
//	@Summary		Shows the current version of the api
//	@Description	shows the current version of the api
//	@Tags			printers
//	@Success		200	{string} string "version #"
//	@Failure		404	{object}	response.APIError	"Printer not found"
//	@Router			/ancillary/version/api [get]
func GetVersionApi(c *fiber.Ctx) error {
	//TODO:  get that version from an env file.
	version := "0.0.1-pre-alpha"
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"version": version})
}
