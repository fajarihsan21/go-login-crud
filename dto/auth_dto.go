package dto

type AuthReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthRes struct {
	Token string `json:"token"`
}