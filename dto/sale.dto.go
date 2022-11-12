package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
	"time"
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
	Data *model.Sale `json:"data"`
}

type UpdateSaleRes struct {
	BaseRes
	Data *model.Sale `json:"data"`
}

type DeleteSaleRes struct {
	BaseRes
}

type SaleBody struct {
	Name      string    `json:"name"`
	Discount  int16     `json:"discount"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	IsApplied bool      `json:"is_applied"`
	Code      string    `json:"code"`
	CreatedBy uuid.UUID `json:"created_by"`
}
