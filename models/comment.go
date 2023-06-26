package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	Id        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	Deskripsi string    `gorm:"size:255;not null" json:"deskripsi"`
	IdUser    uint32    `gorm:"not null;index" json:"idUser"`
	// User        User      `gorm:"foreignKey:IdUser"`
	IdCommunity uint32 `gorm:"not null;index" json:"idCommunity"`
}

// Create
func (c *Comment) CreatedComment(db *gorm.DB) (*Comment, error) {
	err := db.Create(&c).Error
	if err != nil {
		return &Comment{}, err
	}
	return c, nil
}

// delete comment
func (c *Comment) DeleteComment(db *gorm.DB, id uint32) (*Comment, error) {
	err := db.Where("id = ?", id).Delete(&c).Error
	if err != nil {
		return &Comment{}, err
	}
	return c, nil
}
