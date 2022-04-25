package models

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

type UpReq struct {
	F_name       string `json:"f_name"`
	L_name       string `json:"l_name"`
	New_password string `json:"new_password"`
	Old_password string `json:"old_password"`
	Position     string `json:"position"`
}

type Status struct {
	Position string `json:"position"`
}

type CreateUser struct {
	F_name   string `json:"f_name"`
	L_name   string `json:"l_name"`
	Password string `json:"password"`
	Monthly  string `json:"monthly"`
	Position string `json:"position"`
}

type Get struct {
	Id         string `json:"id"`
	F_name     string `json:"f_name"`
	L_name     string `json:"l_name"`
	Password   string `json:"password"`
	Monthly    string `json:"monthly"`
	Position   string `json:"position"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
	Hours      []Hour `json:"hours"`
}

type Hour struct {
	ID        string `json:"id"`
	User_id   string `json:"user_id"`
	Last_name string `json:"last_name"`
	Opened    string `json:"Opened"`
	Daily     string `json:"daily"`
	Monthly   string `json:"monthly"`
	Date      string `json:"date"`
	Klozed    string `json:"klozed"`
}

type GetUser struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type PasswordReq struct {
	Password string `json:"password"`
}

type ById struct {
	Userid string `json:"userid"`
}

type ListResp struct {
	Users []User `json:"users"`
	Count string `json:"count"`
}

type ListReq struct {
	Page  int64 `json:"page" `
	Limit int64 `json:"limit"`
}

type EmptyResp struct {
	Message string `json:"message"`
}
