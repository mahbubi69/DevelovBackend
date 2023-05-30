package models

import "github.com/jinzhu/gorm"

type FeedBack struct {
	Id         uint32 `gorm:"primary_key;auto_increment" json:"id"`
	TepatWaktu bool   `gorm:"default:false;not null" json:"tepatWaktu"`
	Rating     int    `gorm:"size:2;null" json:"rating"`
	Deskripsi  string `gorm:"type:text;null" json:"deskripsi"`
	IdMentor   uint32 `gorm:"not null;index" json:"idMentor"`
	IdUser     uint32 `gorm:"not null;index" json:"idUser"`
}

func (fb *FeedBack) CreatedFeedBack(db *gorm.DB) (*FeedBack, error) {
	err := db.Create(&fb).Error
	if err != nil {
		return &FeedBack{}, err
	}
	return fb, nil
}
