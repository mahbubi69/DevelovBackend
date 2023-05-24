package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"lify_backend/auth"
	"log"
	"os/exec"
	"strconv"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Uuid     string `gorm:"size:100;not null;unique" json:"uuId"`
	Nama     string `gorm:"size:100;not null;unique" json:"nama"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Token    string `gorm:"type:text" json:"token"`
	TglLahir string `gorm:"size:100;not null;" json:"tgl_lahir"`
}

type ResponseUserMapping struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	TglLahir string `json:"tgl_lahir"`
}

func NewUUID() string {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(newUUID)
}

func HashPasswordToSha256(password string) string {
	sum := sha256.Sum256([]byte(password))
	hashedPassword := hex.EncodeToString(sum[:])
	return hashedPassword
}

// query Login
func (u *User) SignIn(db *gorm.DB, id uint32, email, pasword string) (string, error) {
	user := User{}
	if err := db.Where("email = ?", email).Take(&user).Error; err != nil {
		return "", err
	}
	//checking password to sha 256
	hashedPassword256 := HashPasswordToSha256(pasword)
	if hashedPassword256 != user.Password {
		return "", errors.New("Password Salah")
	}
	return auth.CreateToken(id)
}

// update token on login
func (u *User) UpdateTokenUser(db *gorm.DB, email, token string) (string, error) {
	user := User{}
	err := db.Model(&user).
		Where("email = ?", email).
		Update("token", token).Error

	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *User) GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	user := User{}
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// Create
func (u *User) CreateUser(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, err
}

// Read id
func (u *User) GetUserById(db *gorm.DB, id uint32) (*User, error) {
	user := User{}
	err := db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// get all user maping
func (u *User) GetAllUserMap(db *gorm.DB, users []User) []ResponseUserMapping {
	var allUser []ResponseUserMapping
	for _, user := range users {
		var response ResponseUserMapping
		getUser, _ := u.GetUserById(db, uint32(user.Id))
		response.Email = getUser.Email
		response.Nama = getUser.Nama

		allUser = append(allUser, response)
	}

	return allUser
}

// Read All
func (u *User) GetAllUser(db *gorm.DB, pages, offests string) (*[]User, uint64, error) {
	user := []User{}
	var itemCount uint64

	//string to uint 16
	page, _ := strconv.ParseUint(pages, 10, 32)
	offset, _ := strconv.ParseUint(offests, 10, 32)

	err := db.
		Offset((page - 1) * offset).
		Limit(offset).
		Find(&user).
		Count(&itemCount).
		Error
	if err != nil {
		return &[]User{}, 0, err
	}
	return &user, itemCount, nil
}
