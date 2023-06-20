package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

type Mentor struct {
	Id            uint32        `gorm:"primary_key;auto_increment" json:"id"`
	Profile       string        `gorm:"type:text;null" json:"profile"`
	Nama          string        `gorm:"size:100;not null" json:"nama"`
	Email         string        `gorm:"size:100;not null;unique" json:"email"`
	Specialist    string        `gorm:"size:100;not null" json:"specialist"`
	Portofolio    string        `gorm:"type:text;null" json:"portofolio"`
	Salary        string        `gorm:"type:text;null" json:"salary"`
	NamaRekening  string        `gorm:"type:text;null" json:"namaRekening"`
	NomorRekening string        `gorm:"type:text;null" json:"nomorRekening"`
	FeedBack      []FeedBack    `gorm:"Foreignkey:IdMentor;association_foreignkey:Id;" json:"feedBack"`
	MentorTools   []MentorTools `gorm:"Foreignkey:IdMentor;association_foreignkey:Id;" json:"mentorTools"`
}

// Create
func (m *Mentor) CreatedMentor(db *gorm.DB) (*Mentor, error) {
	err := db.Create(&m).Error
	if err != nil {
		return &Mentor{}, err
	}
	return m, nil
}

// Read All
func (m *Mentor) GetAllMentor(db *gorm.DB, pages, offests string) (*[]Mentor, uint64, error) {
	mentor := []Mentor{}
	var itemCount uint64

	//string to uint 16
	page, _ := strconv.ParseUint(pages, 10, 32)
	offset, _ := strconv.ParseUint(offests, 10, 32)

	err := db.Model(mentor).
		Preload("FeedBack").
		Preload("MentorTools").
		Preload("MentorTools.Tools").
		Offset((page - 1) * offset).
		Limit(offset).
		Find(&mentor).
		Count(&itemCount).
		Error
	if err != nil {
		return &[]Mentor{}, 0, err
	}
	return &mentor, itemCount, nil
}

// update
func (m *Mentor) UpdateMentor(db *gorm.DB, id uint32) (*Mentor, error) {
	// mentor := Mentor{}
	err := db.Model(&Mentor{}).Where("id = ?", id).Update(&m).Error
	if err != nil {
		return &Mentor{}, err
	}
	return m, nil
}

// delete mentor
func (m *Mentor) DeleteMentor(db *gorm.DB, id uint32) (*Mentor, error) {
	err := db.Where("id = ?", id).Delete(&m).Error
	if err != nil {
		return &Mentor{}, err
	}
	return m, nil
}
