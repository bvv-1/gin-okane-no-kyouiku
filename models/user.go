package models

import (
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
	user.BeforeSave()
	if err := db.Debug().Model(&User{}).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeSave() error {
	// Ref) https://gorm.io/docs/hooks.html
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil
}
