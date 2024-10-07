package dto

type UserData struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserRegist struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type RegistRes struct {
	Message string `json:"Message"`
	Code    string `json:"Code"`
	Result  string `json:"Result"`
}

type UserReq struct {
	Page  int `json:"Page"`
	Limit int `json:"Limit"`
}

type DelReq struct {
	Id string `json:"id"`
}