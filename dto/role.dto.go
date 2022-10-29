package dto

import (
	"eCommerce/model"
)

type GetAllRoleDataRes struct {
	Total int16        `json:"total"`
	Page  int16        `json:"page"`
	Data  []model.Role `json:"data"`
}

type GetAllRoleRes struct {
	BaseRes
	Data GetAllRoleDataRes `json:"data"`
}

type GetOneRoleRes struct {
	BaseRes
	Data model.Role `json:"data`
}

type UpdateRoleRes struct {
	BaseRes
	Data model.Role `json:"data`
}

type DeleteRoleRes struct {
	BaseRes
}

type RoleBody struct {
	//Name       string    `json:"name"`
	//Category   uuid.UUID `json:"category"`
	//Detail     string    `json:"detail"`
	//RoleManual string    `json:"user_manual"`
	//Price      int16     `json:"price"`
	//CreatedBy  uuid.UUID `json:"created_by"`
}
