package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

type GetAllSaleItemDataRes struct {
	Total int16            `json:"total"`
	Page  int16            `json:"page"`
	Data  []model.SaleItem `json:"data"`
}

type GetAllSaleItemRes struct {
	BaseRes
	Data GetAllSaleItemDataRes `json:"data"`
}

type GetOneSaleItemRes struct {
	BaseRes
	Data *model.SaleItem `json:"data"`
}

type UpdateSaleItemRes struct {
	BaseRes
	Data *model.SaleItem `json:"data"`
}

type DeleteSaleItemRes struct {
	BaseRes
}

type SaleItemBody struct {
	Item   uuid.UUID `json:"item"`
	Amount int16     `json:"amount"`
	Sale   uuid.UUID `json:"sale"`
}
