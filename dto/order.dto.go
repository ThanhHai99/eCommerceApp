package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

type GetAllOrderDataRes struct {
	Total int16         `json:"total"`
	Page  int16         `json:"page"`
	Data  []model.Order `json:"data"`
}

type GetAllOrderRes struct {
	BaseRes
	Data GetAllOrderDataRes `json:"data"`
}

type GetOneOrderRes struct {
	BaseRes
	Data *model.Order `json:"data"`
}

type UpdateOrderRes struct {
	BaseRes
	Data *model.Order `json:"data"`
}

type DeleteOrderRes struct {
	BaseRes
}

type OrderBody struct {
	DeliveryAddress string    `json:"delivery_address"`
	PaymentMethod   string    `json:"payment_method"`
	Paid            bool      `json:"paid"`
	Cost            int16     `json:"cost"`
	CreatedBy       uuid.UUID `json:"created_by"`
}
