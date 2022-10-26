package models

// import (
// 	validation "github.com/go-ozzo/ozzo-validation/v3"
// )

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	City       string `json:"city"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
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

// func (rum *User) Validate() error {
// 	return validation.ValidateStruct(
// 		rum,
// 		validation.Field(&rum.Name, validation.Required, validation.Length(1, 30)),
// 		validation.Field(&rum.City, validation.Required, validation.Length(1, 30)),
// 		validation.Field(&rum.Username, validation.Required, validation.Length(5, 30)),
// 	)
// }
