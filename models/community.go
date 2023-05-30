package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type Community struct {
	Id        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	Title     string     `gorm:"size:255;not null" json:"title"`
	Deskripsi string     `gorm:"size:255;not null" json:"deskripsi"`
	Image     string     `gorm:"type:text;null" json:"image"`
	Comment   []*Comment `gorm:"Foreignkey:IdCommunity;association_foreignkey:Id;" json:"comment"`
}

// ceated
func (c *Community) CreatedCommunity(db *gorm.DB) (*Community, error) {
	err := db.Create(&c).Error
	if err != nil {
		return &Community{}, err
	}
	return c, nil
}

// Read All
func (c *Community) GetACommunity(db *gorm.DB, pages, offests string) (*[]Community, uint64, error) {
	community := []Community{}
	var itemCount uint64

	//string to uint 16
	page, _ := strconv.ParseUint(pages, 10, 32)
	offset, _ := strconv.ParseUint(offests, 10, 32)

	err := db.Model(community).
		Offset((page - 1) * offset).
		Preload("Comment").
		// Preload("Comment.User").
		Limit(offset).
		Find(&community).
		Count(&itemCount).
		Error
	if err != nil {
		return &[]Community{}, 0, err
	}
	return &community, itemCount, nil
}
