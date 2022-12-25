package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
	"time"
)

type GetAllSaleLogDataRes struct {
	BaseRes
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

type CreateSaleLogBody struct {
	Name       string    `json:"name"`
	SaleItemID uuid.UUID `json:"sale_item_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	Amount     int16     `json:"amount"`
	Discount   int16     `json:"discount"`
	IsApplied  bool      `json:"is_applied"`
	Code       string    `json:"code"`
	SaleID     uuid.UUID `json:"sale_id"`
}
