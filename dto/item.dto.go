package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

type GetAllItemDataRes struct {
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
	Data model.Item `json:"data`
}

type UpdateItemRes struct {
	BaseRes
	Data model.Item `json:"data`
}

type DeleteItemRes struct {
	BaseRes
}

type ItemBody struct {
	Name       string    `json:"name"`
	Category   uuid.UUID `json:"category"`
	Detail     string    `json:"detail"`
	UserManual string    `json:"user_manual"`
	Price      int16     `json:"price"`
	CreatedBy  uuid.UUID `json:"created_by"`
}
