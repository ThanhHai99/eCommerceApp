package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
	"time"
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
	Data *model.Warehouse `json:"data"`
}

type UpdateWarehouseRes struct {
	BaseRes
	Data *model.Warehouse `json:"data"`
}

type DeleteWarehouseRes struct {
	BaseRes
}

type WarehouseBody struct {
	ExpirationDate []time.Time `json:"expiration_date"`
	Amount         []int16     `json:"amount"`
	Price          []int16     `json:"price"`
	Item           []uuid.UUID `json:"item"`
	CreatedBy      uuid.UUID   `json:"created_by"`
}
