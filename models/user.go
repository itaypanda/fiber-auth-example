package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid,unique"`
	Name     string    `json:"name" validate:"required,min=3,max=16,alpha"`
	Email    string    `json:"email" gorm:"unique" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=8,max=256"`
}

type SigninUser struct {
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=8,max=256"`
}

var validate *validator.Validate

func ValidateSignupUser(u *User) error {
	validate = validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(u)
}

func ValidateSigninUser(u *SigninUser) error {
	validate = validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(u)
}

