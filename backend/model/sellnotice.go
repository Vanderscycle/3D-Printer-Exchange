package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SellNotice struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;"`
	User	User `gorm:"references:ID"`
	Printer  Printer `gorm:"references:ID"`
	Issue string `json:"issue"`
	Info string `json:"info"`
	Status string `json:"status"`
	Location string `json:"location"` // TODO figure out geolocation

}

func (sellNotice *SellNotice) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	sellNotice.ID = uuid.New()
	return
}
