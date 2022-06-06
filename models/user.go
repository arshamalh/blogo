package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password []byte
	Email    string
	Posts    []Post `gorm:"foreignKey:AuthorID"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
