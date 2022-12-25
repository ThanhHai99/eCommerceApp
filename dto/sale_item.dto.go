package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

type GetAllSaleItemDataRes struct {
	BaseRes
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

type CreateSaleItemBody struct {
	ItemID uuid.UUID `json:"item_id"`
	Amount int16     `json:"amount"`
	SaleID uuid.UUID `json:"sale_id"`
}

type UpdateSaleItemBody struct {
	CreateSaleItemBody
}
