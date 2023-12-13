package models

import (
	"github.com/arshamalh/blogo/log"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            uint `bun:"id"`
	bun.BaseModel `bun:"user"`
	Username      string     `json:"username" bun:"unique"`
	Password      []byte     `json:"-"`
	Email         string     `json:"email"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	Posts         []*Post    `json:"posts" bun:"rel:has-many"`
	Comments      []*Comment `json:"comments" bun:"rel:has-many"`
	Role          *Role      `json:"role" bun:"rel:belongs-to"`
	RoleID        uint       `json:"role_id"`
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Gl.Error(err.Error())
		return err
	}

	user.Password = hashedPassword
	return nil
}

func (user *User) ComparePasswords(password string) error {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		log.Gl.Error(err.Error())
		return err
	}
	return nil
}
