package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

type GetAllItemOrderDataRes struct {
	BaseRes
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

type CreateItemOrderBody struct {
	ItemID  uuid.UUID `json:"item_id"`
	Amount  int16     `json:"amount"`
	OrderID uuid.UUID `json:"order_id"`
}
