package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

type GetAllUserDataRes struct {
	Total int16        `json:"total"`
	Page  int16        `json:"page"`
	Data  []model.User `json:"data"`
}

type GetAllUserRes struct {
	BaseRes
	Data GetAllUserDataRes `json:"data"`
}

type GetOneUserRes struct {
	BaseRes
	Data *model.User `json:"data"`
}

type UpdateUserRes struct {
	BaseRes
	Data *model.User `json:"data"`
}

type DeleteUserRes struct {
	BaseRes
}

type UserBody struct {
	UserName    string    `json:"user_name"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	IsLocked    bool      `json:"is_locked"`
	Role        uuid.UUID `json:"role"`
	VerifyToken string    `json:"verify_token"`
	IsActive    bool      `json:"is_active"`
}
