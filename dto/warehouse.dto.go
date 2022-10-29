package dto

import (
	"eCommerce/model"
)

type GetAllWarehouseDataRes struct {
	Total int16             `json:"total"`
	Page  int16             `json:"page"`
	Data  []model.Warehouse `json:"data"`
}

type GetAllWarehouseRes struct {
	BaseRes
	Data GetAllWarehouseDataRes `json:"data"`
}

type GetOneWarehouseRes struct {
	BaseRes
	Data model.Warehouse `json:"data`
}

type UpdateWarehouseRes struct {
	BaseRes
	Data model.Warehouse `json:"data`
}

type DeleteWarehouseRes struct {
	BaseRes
}

type WarehouseBody struct {
	//Name       string    `json:"name"`
	//Category   uuid.UUID `json:"category"`
	//Detail     string    `json:"detail"`
	//WarehouseManual string    `json:"user_manual"`
	//Price      int16     `json:"price"`
	//CreatedBy  uuid.UUID `json:"created_by"`
}
