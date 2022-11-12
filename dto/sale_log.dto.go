package dto

import (
	"eCommerce/model"
)

type GetAllSaleLogDataRes struct {
	Total int16           `json:"total"`
	Page  int16           `json:"page"`
	Data  []model.SaleLog `json:"data"`
}

type GetAllSaleLogRes struct {
	BaseRes
	Data GetAllSaleLogDataRes `json:"data"`
}

type GetOneSaleLogRes struct {
	BaseRes
	Data *model.SaleLog `json:"data"`
}

type UpdateSaleLogRes struct {
	BaseRes
	Data *model.SaleLog `json:"data"`
}

type DeleteSaleLogRes struct {
	BaseRes
}

type SaleLogBody struct {
	//Name       string    `json:"name"`
	//Category   uuid.UUID `json:"category"`
	//Detail     string    `json:"detail"`
	//SaleLogManual string    `json:"user_manual"`
	//Price      int16     `json:"price"`
	//CreatedBy  uuid.UUID `json:"created_by"`
}
