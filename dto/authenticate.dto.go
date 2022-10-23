package dto

type RegisterBody struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterRes struct {
	BaseRes
	IsSentMailActive bool `json:"is_sent_mail_active"`
}

type LoginBody struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginRes struct {
	BaseRes
	AccessToken string `json:"access_token"`
}

type VerifyRes struct {
	BaseRes
}
