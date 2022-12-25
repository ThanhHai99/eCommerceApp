package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
	"time"
)

type GetAllWarehouseLogDataRes struct {
	BaseRes
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
	Data *model.WarehouseLog `json:"data"`
}

type UpdateWarehouseLogRes struct {
	BaseRes
	Data *model.WarehouseLog `json:"data"`
}

type DeleteWarehouseLogRes struct {
	BaseRes
}

type CreateWarehouseLogBody struct {
	Status         string    `json:"status"`
	Price          int16     `json:"price"`
	ExpirationDate time.Time `json:"expiration_date"`
	Amount         int16     `json:"amount"`
	WarehouseID    uuid.UUID `json:"warehouse_id"`
	ItemID         uuid.UUID `json:"item_id"`
}
