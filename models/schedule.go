package models

import "github.com/jinzhu/gorm"

type Schedule struct {
	Id       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Hari     string `gorm:"size:100;null" json:"hari"`
	Jam      string `gorm:"size:100;null" json:"Jam"`
	IdUser   uint32 `gorm:"not null;index" json:"idUser"`
	IdMentor uint32 `gorm:"not null;index" json:"idMentor"`
}

// ceated
func (s *Schedule) CreatedSchedule(db *gorm.DB) (*Schedule, error) {
	err := db.Create(&s).Error
	if err != nil {
		return &Schedule{}, err
	}
	return s, nil
}

// Read
func (s *Schedule) GetSchedule(db *gorm.DB) (*[]Schedule, uint64, error) {
	schedule := []Schedule{}
	var itemCount uint64

	err := db.Find(&schedule).
		Count(&itemCount).Error
	if err != nil {
		return &[]Schedule{}, 0, err
	}
	return &schedule, itemCount, nil
}

// update
func (s *Schedule) UpdateSchedule(db *gorm.DB, id uint32) (*Schedule, error) {

	err := db.Model(&Schedule{}).Where("id = ?", id).Update(&s).Error

	if err != nil {
		return &Schedule{}, err
	}

	return s, nil

}
