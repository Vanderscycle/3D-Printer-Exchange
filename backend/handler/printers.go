package handler

import (
	"github.com/Vanderscycle/3D-Printer-Exchange/database"
	"github.com/Vanderscycle/3D-Printer-Exchange/model"
	_ "github.com/Vanderscycle/3D-Printer-Exchange/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetAllPrinters	godoc
//
//	@Summary		Show all printers listed
//	@Description	show all printers
//	@Tags			printers
//	@Success		200	{array}		model.Printer
//	@Failure		404	{object}	response.APIError	"Printers not found"
//	@Router			/api/printer [get]
func GetAllPrinters(c *fiber.Ctx) error {
	//TODO add a get all for a user
	db := database.DB.Db
	var printers []model.Printer

	// find all printers in the database
	db.Find(&printers)

	// If no printer found, return an error
	if len(printers) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Printer not found", "data": nil})
	}

	// return printers
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Printers Found", "data": printers})
}

// GetSinglePrinter 	godoc
//
//	@Summary		show the data for printer
//	@Description	get printer data by ID
//	@Tags			printers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Printer Id"
//	@Success		200	{object}	model.Printer
//	@Failure		404	{object}	response.APIError	"Printer not found"
//	@Router			/api/printer/{id} [get]
func GetSinglePrinter(c *fiber.Ctx) error {
	db := database.DB.Db

	// get id params
	id := c.Params("id")

	var printer model.Printer

	// find single printer in the database by id
	db.Find(&printer, "id = ?", id)

	if printer.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Printer not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Printer Found", "data": printer})
}

// CreatePrinter godoc
//
//	@Summary		add a new printer to the DB
//	@Description	create a new printer object
//	@Tags			printers
//	@Accept			json
//	@Produce		json
//	@Param			message	body		model.Printer	true	"Printer Data"
//	@Success		200		{object}	model.Printer
//	@Failure		500		{object}	response.APIError	"Error with input"
//	@Router			/api/printer [post]
func CreatePrinter(c *fiber.Ctx) error {
	db := database.DB.Db
	printer := new(model.Printer)

	// Store the body in the printer and return error if encountered
	if err := c.BodyParser(printer); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err, "input": printer})
	}

	err := db.Create(&printer).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not add printer", "data": err})
	}

	// Return the created printer data
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "This printer has been added to your account", "data": printer})
}

// UpdatePrinter 	godoc
//
//	@Summary		Updates a user's printer info
//	@Description	Updates a printer data
//	@Tags			printers
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Printer Id"
//	@Param			message	body		model.Printer		true	"Printer Data"
//	@Success		200		{object}	model.Printer		"Printer updated"
//	@Failure		500		{object}	response.APIError	"Error with input"
//	@Failure		404		{object}	response.APIError	"Printer not found"
//	@Router			/api/printer/{id} [patch]
func UpdatePrinter(c *fiber.Ctx) error {
	type UpdatePrinter struct {
		Brand       string `gorm:"brand"`
		Name        string `json:"name"`
		PowerSupply string `json:"powerSupply"`
		Probe       string `json:"probe"`
		Board       string `json:"board"`
		Hotend      string `json:"hotend"`
		Extruder    string `json:"extruder"`
		Nozzle      string `json:"nozzle"`
		BuildVolume string `json:"buildVolume"`
		Mods        string `json:"mods"`
	}

	db := database.DB.Db

	var printer model.Printer

	// get id params
	id := c.Params("id")

	// find single user in the database by id
	db.Find(&printer, "id = ?", id)

	if printer.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Printer not found", "data": nil})
	}

	var updatePrinterData UpdatePrinter
	err := c.BodyParser(&updatePrinterData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	//REVIEW: Is there an easier way?
	printer.Brand = updatePrinterData.Brand
	printer.Name = updatePrinterData.Name
	printer.PowerSupply = updatePrinterData.PowerSupply
	printer.Probe = updatePrinterData.Probe
	printer.Board = updatePrinterData.Board
	printer.Hotend = updatePrinterData.Hotend
	printer.Extruder = updatePrinterData.Extruder
	printer.Nozzle = updatePrinterData.Nozzle
	printer.BuildVolume = updatePrinterData.BuildVolume
	printer.Mods = updatePrinterData.Mods

	// Save the Changes
	db.Save(&printer)

	// Return the updated printer data
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "printer updated", "data": printer})

}

// DeletePrinterByID godoc
//
//	@Summary		Delete a user's printer
//	@Description	Delete a printer
//	@Tags			printers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int					true	"Printer Id"
//	@Success		200	{object}	model.Printer		"Printer deleted"
//	@Failure		404	{object}	response.APIError	"Printer not found"
//	@Router			/api/printer/{id} [delete]
func DeletePrinterByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var printer model.Printer

	// get id params
	id := c.Params("id")

	// find single printer in the database by id
	db.Find(&printer, "id = ?", id)

	if printer.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Printer not found", "data": nil})

	}

	err := db.Delete(&printer, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete printer data", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Printer deleted"})
}
