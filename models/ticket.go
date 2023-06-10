package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Ticket struct {
	Id         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CodeTiket  string    `gorm:"size:100;not null;" json:"struk"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	IdSchedule uint32    `gorm:"foreignKey:id_schedule;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"idSchedule"`
	// IdSchedule uint32    `gorm:"foreignKey:id_schedule;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"idSchedule"`
}

// created
func (tk *Ticket) CreatedTicket(db *gorm.DB) (*Ticket, error) {
	err := db.Create(&tk).Error
	if err != nil {
		return &Ticket{}, err
	}
	return tk, nil
}

// update
func (tk *Ticket) UpdateTicket(db *gorm.DB, id uint32) (*Ticket, error) {
	// mentor := Mentor{}
	err := db.Model(&Ticket{}).Where("id = ?", id).Update(&tk).Error
	if err != nil {
		return &Ticket{}, err
	}
	return tk, nil
}
