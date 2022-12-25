package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

type GetAllPriceLogDataRes struct {
	BaseRes
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

type CreatePriceLogBody struct {
	ItemID uuid.UUID `json:"item_id"`
	Price  int16     `json:"price"`
}
