package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

type GetAllItemDataRes struct {
	BaseRes
	Total int16        `json:"total"`
	Page  int16        `json:"page"`
	Data  []model.Item `json:"data"`
}

type GetAllItemRes struct {
	BaseRes
	Data GetAllItemDataRes `json:"data"`
}

type GetOneItemRes struct {
	BaseRes
	Data *model.Item `json:"data"`
}

type UpdateItemRes struct {
	BaseRes
	Data *model.Item `json:"data"`
}

type DeleteItemRes struct {
	BaseRes
}

type CreateItemBody struct {
	Name       string    `json:"name"`
	CategoryID uuid.UUID `json:"category_id"`
	Detail     string    `json:"detail"`
	UserManual string    `json:"user_manual"`
	Price      int16     `json:"price"`
}

type UpdateItemBody struct {
	CreateItemBody
}
