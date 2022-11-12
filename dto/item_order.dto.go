package dto

import (
	"eCommerce/model"
)

type GetAllItemOrderDataRes struct {
	Total int16             `json:"total"`
	Page  int16             `json:"page"`
	Data  []model.ItemOrder `json:"data"`
}

type GetAllItemOrderRes struct {
	BaseRes
	Data GetAllItemOrderDataRes `json:"data"`
}

type GetOneItemOrderRes struct {
	BaseRes
	Data *model.ItemOrder `json:"data"`
}

type UpdateItemOrderRes struct {
	BaseRes
	Data *model.ItemOrder `json:"data"`
}

type DeleteItemOrderRes struct {
	BaseRes
}

type ItemOrderBody struct {
	//Name       string    `json:"name"`
	//Category   uuid.UUID `json:"category"`
	//Detail     string    `json:"detail"`
	//ItemOrderManual string    `json:"user_manual"`
	//Price      int16     `json:"price"`
	//CreatedBy  uuid.UUID `json:"created_by"`
}
