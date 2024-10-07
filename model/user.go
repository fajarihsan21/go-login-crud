package model

import Mstring "go-login-crud/model/string"

type User struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type RespUser struct {
	UserId   Mstring.NullString `json:"user_id"`
	Username Mstring.NullString `json:"username"`
	Password Mstring.NullString `json:"password"`
	Email    Mstring.NullString `json:"email"`
	Phone    Mstring.NullString `json:"phone"`
}

func (u RespUser) CreateUserData() User {
	resp := User{
		UserId: u.UserId.String,
		Username: u.Username.String,
		Password: u.Password.String,
		Email: u.Email.String,
		Phone: u.Phone.String,
	}

	return resp
}

type UserRes struct {
	Message string `json:"Message"`
	Code    string `json:"Code"`
	Result  User `json:"Result"`
}

type UsersRes struct {
	Message string `json:"Message"`
	Code    string `json:"Code"`
	Result  []User `json:"Result"`
}