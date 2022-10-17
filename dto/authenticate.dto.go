package dto

type RegisterBody struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterRes struct {
	BaseRes
}

type LoginBody struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginRes struct {
	BaseRes
}

type VerifyRes struct {
	BaseRes
}
