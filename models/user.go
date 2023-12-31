package models

import (
	"gin-okane-no-kyouiku/utils/token"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

func InsertUser(db *gorm.DB, email string, password string) error {
	user := User{
		Email:    email,
		Password: password,
	}
	if err := db.Debug().Model(&User{}).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeSave(db *gorm.DB) error {
	// Ref) https://gorm.io/docs/hooks.html
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(db *gorm.DB, email string, password string) (string, error) {
	var user User
	if err := db.Debug().Model(&User{}).Where("email = ?", email).Take(&user).Error; err != nil {
		return "", err
	}
	if err := VerifyPassword(user.Password, password); err != nil {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
