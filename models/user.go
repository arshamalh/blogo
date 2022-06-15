package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"uniqueIndex"`
	Password  []byte `json:"-"`
	Email     string `json:"email"`
	FisrtName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Posts     []Post `gorm:"foreignKey:AuthorID"`
	Role      Role   `gorm:"foreignKey:RoleID"`
	RoleID    uint   `json:"role_id"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
