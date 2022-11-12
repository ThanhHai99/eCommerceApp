package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
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
	Name       string    `json:"name"`
	Detail     string    `json:"detail"`
	UserManual string    `json:"user_manual"`
	CreatedBy  uuid.UUID `json:"created_by"`
	Category   uuid.UUID `json:"category"`
	Item       uuid.UUID `json:"item"`
}
