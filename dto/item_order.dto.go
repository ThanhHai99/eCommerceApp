package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
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
	Item   uuid.UUID `json:"item"`
	Amount int16     `json:"amount"`
	Order  uuid.UUID `json:"order"`
}
