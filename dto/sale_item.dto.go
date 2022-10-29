package dto

import (
	"eCommerce/model"
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
	Data model.SaleItem `json:"data`
}

type UpdateSaleItemRes struct {
	BaseRes
	Data model.SaleItem `json:"data`
}

type DeleteSaleItemRes struct {
	BaseRes
}

type SaleItemBody struct {
	//Name       string    `json:"name"`
	//Category   uuid.UUID `json:"category"`
	//Detail     string    `json:"detail"`
	//SaleItemManual string    `json:"user_manual"`
	//Price      int16     `json:"price"`
	//CreatedBy  uuid.UUID `json:"created_by"`
}
