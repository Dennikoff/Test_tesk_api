package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                uint   `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
	NickName          string `json:"nickname"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, is.Email, validation.Required),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 100)),
		validation.Field(&u.NickName, validation.Required),
	)
}

func (u *User) BeforeCreate() error {
	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 5)
	if err != nil {
		return err
	}
	u.EncryptedPassword = string(pass)
	u.Password = ""
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
}
