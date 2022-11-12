package dto

import (
	"eCommerce/model"
)

type GetAllInvoiceDataRes struct {
	Total int16        `json:"total"`
	Page  int16        `json:"page"`
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
	//Name       string    `json:"name"`
	//Category   uuid.UUID `json:"category"`
	//Detail     string    `json:"detail"`
	//UserManual string    `json:"user_manual"`
	//Price      int16     `json:"price"`
	//CreatedBy  uuid.UUID `json:"created_by"`
}
