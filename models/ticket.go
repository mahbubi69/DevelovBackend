package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Ticket struct {
	Id         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CodeTiket  string    `gorm:"size:100;not null;" json:"struk"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	IdUser     uint32    `gorm:"not null;index" json:"idUser"`
	IdSchedule uint32    `gorm:"null;lforeignKey:id_schedule;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"idSchedule"`
}

// created
func (tk *Ticket) CreatedTicket(db *gorm.DB) (*Ticket, error) {
	err := db.Model(&Ticket{}).Create(&tk).Error
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
