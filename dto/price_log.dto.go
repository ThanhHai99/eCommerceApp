package dto

import (
	"eCommerce/model"
)

type GetAllPriceLogDataRes struct {
	Total int16            `json:"total"`
	Page  int16            `json:"page"`
	Data  []model.PriceLog `json:"data"`
}

type GetAllPriceLogRes struct {
	BaseRes
	Data GetAllPriceLogDataRes `json:"data"`
}

type GetOnePriceLogRes struct {
	BaseRes
	Data *model.PriceLog `json:"data"`
}

type UpdatePriceLogRes struct {
	BaseRes
	Data *model.PriceLog `json:"data"`
}

type DeletePriceLogRes struct {
	BaseRes
}

type PriceLogBody struct {
	//Name       string    `json:"name"`
	//Category   uuid.UUID `json:"category"`
	//Detail     string    `json:"detail"`
	//PriceLogManual string    `json:"user_manual"`
	//Price      int16     `json:"price"`
	//CreatedBy  uuid.UUID `json:"created_by"`
}
