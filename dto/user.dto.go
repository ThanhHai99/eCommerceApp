package dto

import (
	"eCommerce/model"
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
	Data model.User `json:"data`
}

type UpdateUserRes struct {
	BaseRes
	Data model.User `json:"data`
}

type DeleteUserRes struct {
	BaseRes
}

type UserBody struct {
	//Name       string    `json:"name"`
	//Category   uuid.UUID `json:"category"`
	//Detail     string    `json:"detail"`
	//UserManual string    `json:"user_manual"`
	//Price      int16     `json:"price"`
	//CreatedBy  uuid.UUID `json:"created_by"`
}
