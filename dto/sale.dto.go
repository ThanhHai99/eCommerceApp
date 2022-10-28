package dto

import (
	"eCommerce/model"
)

type GetAllSaleDataRes struct {
	Total int16        `json:"total"`
	Page  int16        `json:"page"`
	Data  []model.Sale `json:"data"`
}

type GetAllSaleRes struct {
	BaseRes
	Data GetAllSaleDataRes `json:"data"`
}

type GetOneSaleRes struct {
	BaseRes
	Data model.Sale `json:"data`
}

type UpdateSaleRes struct {
	BaseRes
	Data model.Sale `json:"data`
}

type DeleteSaleRes struct {
	BaseRes
}

type SaleBody struct {
	//Name       string    `json:"name"`
	//Category   uuid.UUID `json:"category"`
	//Detail     string    `json:"detail"`
	//UserManual string    `json:"user_manual"`
	//Price      int16     `json:"price"`
	//CreatedBy  uuid.UUID `json:"created_by"`
}
