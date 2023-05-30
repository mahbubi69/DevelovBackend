package models

import "github.com/jinzhu/gorm"

type MentorTools struct {
	Id       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	IdTools  uint32 `gorm:"not null;index" json:"idTools"`
	IdMentor uint32 `gorm:"not null;index" json:"idMentor"`
	Tools    Tools  `gorm:"foreignKey:IdTools"`
}

// created
func (mt *MentorTools) CreatedMentorTools(db *gorm.DB) (*MentorTools, error) {
	err := db.Create(&mt).Error
	if err != nil {
		return &MentorTools{}, err
	}
	return mt, nil
}
