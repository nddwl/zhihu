// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type InfoRequest struct {
}

type InfoResponse struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type LoginRequest struct {
	Mobile           string `json:"mobile" validate:"required,mobile"`
	VerificationCode string `json:"verification_code" validate:"required,len=6,number"`
}

type LoginResponse struct {
	UserId int64 `json:"user_d"`
	Token  Token `json:"token"`
}

type RegisterRequest struct {
	Name             string `json:"name" validate:"required,name"`
	Mobile           string `json:"mobile" validate:"required,mobile"`
	Password         string `json:"password" validate:"required,password"`
	VerificationCode string `json:"verification_code" validate:"required,len=6,number"`
}

type RegisterResponse struct {
	UserId int64 `json:"user_id"`
	Token  Token `json:"token"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
}

type VerificationRequest struct {
	Mobile string `json:"mobile" validate:"required,mobile"`
}

type VerificationResponse struct {
}
