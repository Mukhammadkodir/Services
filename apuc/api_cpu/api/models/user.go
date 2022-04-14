package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v3"
)

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	City       string `json:"city"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}

type By_Id struct {
	User_id string `json:"user_id"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	City     string `json:"city"`
}

type UpdateUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	City     string `json:"city"`
}

type List_Req struct {
	Page  string `json:"page"`
	Limit string `json:"limit"`
}

func (rum *User) Validate() error {
	return validation.ValidateStruct(
		rum,
		validation.Field(&rum.Name, validation.Required, validation.Length(1, 30)),
		validation.Field(&rum.City, validation.Required, validation.Length(1, 30)),
		validation.Field(&rum.Username, validation.Required, validation.Length(5, 30), validation.Match(regexp.MustCompile("^[0-9a-z_.]+$"))),
	)
}
