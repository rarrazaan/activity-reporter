package dto

type ForgetPWReq struct {
	Email     string `json:"email"`
}

type ForgetPWRes struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}
