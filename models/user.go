package models

import (
	"crypto/sha256"
	"develov_be/auth"
	"encoding/hex"
	"log"
	"os/exec"
	"strconv"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Uuid     string `gorm:"size:100;not null;unique" json:"uuId"`
	UserName string `gorm:"size:100;not null;unique" json:"userName"`
	Profile  string `gorm:"type:text;null" json:"profile"`
	Nama     string `gorm:"size:100;not null" json:"nama"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	NoHp     string `gorm:"size:100;not null;unique" json:"noHp"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Purpose  string `gorm:"type:text;null" json:"purpose"`
	Role     int    `gorm:"size:2" json:"role"`
	Token    string `gorm:"type:text;null" json:"token"`
	Otp      string `gorm:"type:text;null" json:"otp"`
}

type ResponseUserMapping struct {
	Nama     string `json:"nama"`
	Profile  string `json:"profile"`
	UserName string `json:"userName"`
	NoHp     string `json:"noHp"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
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

	if err := db.Where("email = ?", email).
		Take(&user).
		Error; err != nil {
		return "maaf email anda salah", err
	}

	//checking password to sha 256
	hashedPassword256 := HashPasswordToSha256(pasword)
	var err error
	if hashedPassword256 != user.Password {
		return "Maaf password anda salah", err
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

// read
func (u *User) GetUserByToken(db *gorm.DB, token string) (*User, error) {
	user := User{}
	err := db.Where("token = ?", token).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// get by token
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
	return u, nil
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

		response.Nama = getUser.Nama
		response.Profile = getUser.Profile
		response.UserName = getUser.UserName
		response.Email = getUser.Email
		response.NoHp = getUser.NoHp
		response.Role = getUser.Role

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

// update profile
func (u *User) UpdateImage(db *gorm.DB, token, image string) (*User, error) {

	err := db.Model(&User{}).
		Where("token = ?", token).
		Update("profile", image).
		Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

// delete user
func (u *User) DeleteUser(db *gorm.DB, id uint32) (*User, error) {
	err := db.Where("id = ?", id).Delete(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// cek email
func (u *User) CekEmail(db *gorm.DB, email string) (*User, error) {
	user := User{}
	err := db.Where("email = ?", email).
		Take(&u).
		Error

	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

// cek Otp
func (u *User) CekOtp(db *gorm.DB, otp string) (*User, error) {
	user := User{}
	err := db.Where("otp = ?", otp).
		Take(&u).
		Error

	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

// update OTP
func (u *User) UpdateOtp(db *gorm.DB, token, otp string) (*User, error) {
	user := User{}

	err := db.Model(&user).
		Where("token = ?", token).
		Update("otp", otp).
		Error

	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

// update password
func (u *User) UpdatePassword(db *gorm.DB, token, password string) (*User, error) {
	user := User{}

	err := db.Model(&user).
		Where("token = ?", token).
		Update("password", password).
		Error

	if err != nil {
		return &User{}, err
	}

	return &user, nil
}
