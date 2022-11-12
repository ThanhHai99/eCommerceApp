package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
	"time"
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
	Name      string    `json:"name"`
	SaleItem  uuid.UUID `json:"sale_item"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Amount    int16     `json:"amount"`
	Discount  int16     `json:"discount"`
	IsApplied bool      `json:"is_applied"`
	Code      string    `json:"code"`
	CreatedBy uuid.UUID `json:"created_by"`
	Sale      uuid.UUID `json:"sale"`
}
