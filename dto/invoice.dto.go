package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

type GetAllInvoiceDataRes struct {
	Total int16           `json:"total"`
	Page  int16           `json:"page"`
	Data  []model.Invoice `json:"data"`
}

type GetAllInvoiceRes struct {
	BaseRes
	Data GetAllInvoiceDataRes `json:"data"`
}

type GetOneInvoiceRes struct {
	BaseRes
	Data *model.Invoice `json:"data"`
}

type UpdateInvoiceRes struct {
	BaseRes
	Data *model.Invoice `json:"data"`
}

type DeleteInvoiceRes struct {
	BaseRes
}

type InvoiceBody struct {
	Order     int16     `json:"order"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Cost      int16     `json:"cost"`
	CreatedBy uuid.UUID `json:"created_by"`
}
