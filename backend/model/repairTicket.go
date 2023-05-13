package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RepairTicket struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;"`
	User	User `gorm:"references:ID"`
	Printer  Printer `gorm:"references:ID"`
	Issue string `json:"issue"`
	Info string `json:"info"`
	Fix string `json:"fix"` // only repairmen (people tasked with repairs) can make changes
	AdditionalInfo string `json:"additionalinfo"`
	Status string `json:"status"`
	Tracking string `json:"tracking"` //shipping trakcing info
	Estimate string `json:"estimate"`
	Bill string `json:"bill"`

}

func (repairTicket *RepairTicket) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	repairTicket.ID = uuid.New()
	return
}
