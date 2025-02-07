package models

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
	Email    string `json:"email"`

	Orders []Order `json:"orders"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Validate() error {
	v := validator.New()
	return v.Struct(u)
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
