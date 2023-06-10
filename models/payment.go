package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Payment struct {
	Id        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	NoRek     string    `gorm:"size:100;not null;" json:"noRek"`
	Struk     string    `gorm:"size:100;not null;" json:"struk"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	IdUser    uint32    `gorm:"not null;index" json:"idUser"`
}

// created
func (p *Payment) CreatedPayment(db *gorm.DB) (*Payment, error) {
	err := db.Create(&p).Error
	if err != nil {
		return &Payment{}, err
	}
	return p, nil
}

// read
func (p *Payment) GetAllPayment(db *gorm.DB) (*[]Payment, uint64, error) {
	payment := []Payment{}
	var itemCount uint64

	err := db.Find(&payment).
		Count(&itemCount).Error
	if err != nil {
		return &[]Payment{}, 0, err
	}
	return &payment, itemCount, nil
}
