package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
) // User struct

type Printer struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;"`
	User        string    `gorm:"user"`
	Brand       string    `gorm:"brand"`
	Name        string    `json:"name"`
	PowerSupply string    `json:"powerSupply"`
	Probe       string    `json:"probe"`
	Board       string    `json:"board"`
	Hotend      string    `json:"hotend"`
	Extruder    string    `json:"extruder"`
	Nozzle      string    `json:"nozzle"`
	BuildVolume string    `json:"buildVolume"`
	Mods        string    `json:"mods"` //REVIEW: should be a k/v nested obj instead of a str
} // Users struct

type Printers struct {
	Printers []Printer `json:"printer"`
}

func (printer *Printer) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	printer.ID = uuid.New()
	return
}
