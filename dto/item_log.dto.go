package dto

import (
	"eCommerce/model"
)

type GetAllItemLogDataRes struct {
	Total int16           `json:"total"`
	Page  int16           `json:"page"`
	Data  []model.ItemLog `json:"data"`
}

type GetAllItemLogRes struct {
	BaseRes
	Data GetAllItemLogDataRes `json:"data"`
}

type GetOneItemLogRes struct {
	BaseRes
	Data *model.ItemLog `json:"data"`
}

type UpdateItemLogRes struct {
	BaseRes
	Data *model.ItemLog `json:"data"`
}

type DeleteItemLogRes struct {
	BaseRes
}

type ItemLogBody struct {
	//Name       string    `json:"name"`
	//Category   uuid.UUID `json:"category"`
	//Detail     string    `json:"detail"`
	//ItemLogManual string    `json:"user_manual"`
	//Price      int16     `json:"price"`
	//CreatedBy  uuid.UUID `json:"created_by"`
}
