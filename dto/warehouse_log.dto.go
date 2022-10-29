package dto

import (
	"eCommerce/model"
)

type GetAllWarehouseLogDataRes struct {
	Total int16                `json:"total"`
	Page  int16                `json:"page"`
	Data  []model.WarehouseLog `json:"data"`
}

type GetAllWarehouseLogRes struct {
	BaseRes
	Data GetAllWarehouseLogDataRes `json:"data"`
}

type GetOneWarehouseLogRes struct {
	BaseRes
	Data model.WarehouseLog `json:"data`
}

type UpdateWarehouseLogRes struct {
	BaseRes
	Data model.WarehouseLog `json:"data`
}

type DeleteWarehouseLogRes struct {
	BaseRes
}

type WarehouseLogBody struct {
	//Name       string    `json:"name"`
	//Category   uuid.UUID `json:"category"`
	//Detail     string    `json:"detail"`
	//WarehouseLogManual string    `json:"user_manual"`
	//Price      int16     `json:"price"`
	//CreatedBy  uuid.UUID `json:"created_by"`
}
