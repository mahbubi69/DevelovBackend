package models

import "github.com/jinzhu/gorm"

type Tools struct {
	Id   uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Logo string `gorm:"type:text;null" json:"image"`
	Nama string `gorm:"size:100;null" json:"nama"`
}

// created
func (t *Tools) CreatedTools(db *gorm.DB) (*Tools, error) {
	err := db.Create(&t).Error
	if err != nil {
		return &Tools{}, err
	}
	return t, nil
}

// read
func (t *Tools) GetTools(db *gorm.DB) (*[]Tools, uint64, error) {
	tools := []Tools{}
	var itemCount uint64
	err := db.Find(&tools).
		Count(&itemCount).
		Error
	if err != nil {
		return &[]Tools{}, 0, err
	}
	return &tools, itemCount, nil
}

// update
func (t *Tools) UpdateTools(db *gorm.DB, id uint32) (*Tools, error) {
	tools := Tools{}
	err := db.Model(tools).Where("id = ?", id).Update(&tools).Error
	if err != nil {
		return &Tools{}, err
	}
	return &tools, nil
}

// delete tools
func (t *Tools) DeleteTools(db *gorm.DB, id uint32) (*Tools, error) {
	err := db.Where("id = ?", id).Delete(&t).Error
	if err != nil {
		return &Tools{}, err
	}
	return t, nil
}
