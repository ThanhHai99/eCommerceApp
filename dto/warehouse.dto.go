package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
	"time"
)

type GetAllWarehouseDataRes struct {
	BaseRes
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
	Data *model.Warehouse `json:"data"`
}

type UpdateWarehouseRes struct {
	BaseRes
	Data *model.Warehouse `json:"data"`
}

type DeleteWarehouseRes struct {
	BaseRes
}

type CreateExportingBody struct {
	Amount []int16     `json:"amount"`
	ItemID []uuid.UUID `json:"item_id"`
}

type CreateImportingBody struct {
	CreateExportingBody
	ExpirationDate []time.Time `json:"expiration_date"`
	Price          []int16     `json:"price"`
}
