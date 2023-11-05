package models

import (
	"github.com/arshamalh/blogo/log"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `json:"username" gorm:"uniqueIndex"`
	Password  []byte    `json:"-"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Posts     []Post    `gorm:"foreignKey:AuthorID"`
	Comments  []Comment `gorm:"foreignKey:UserID"`
	Role      Role      `gorm:"foreignKey:RoleID"`
	RoleID    uint      `json:"role_id"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Gl.Error("Error hashing password for user", zap.String("username", user.Username), zap.Error(err))
		return
	}

	user.Password = hashedPassword
}

func (user *User) ComparePasswords(password string) error {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		log.Gl.Error("Error: Password comparison failed for user", zap.String("username", user.Username), zap.Error(err))

	}
	return err
}
