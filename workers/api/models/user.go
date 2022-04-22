package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	_"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	Id         string `json:"id"`
	F_name     string `json:"f_name"`
	L_name     string `json:"l_name"`
	Password   string `json:"password"`
	Monthly    string `json:"monthly"`
	Position   string `json:"position"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}

type GetUser struct {
	Id         string `json:"id"`
	F_name     string `json:"f_name"`
	L_name     string `json:"l_name"`
	Password   string `json:"password"`
	Monthly    string `json:"monthly"`
	Position   string `json:"position"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
	Hours      Hour   `json:"hours"`
}

type Hour struct {
	ID        string `json:"id"`
	User_id   string `json:"user_id"`
	Last_name string `json:"last_name"`
	Opened    string `json:"Opened"`
	Closed    string `json:"closed`
	Daily     string `json:"daily"`
	Monthly   string `json:"monthly"`
}



func (rum *User) Validate() error {
	return validation.ValidateStruct(
		rum,
		validation.Field(&rum.Password, validation.Required, validation.Length(8, 15), Match(regexp.MustCompile("[0-9]"))),
	)
}
