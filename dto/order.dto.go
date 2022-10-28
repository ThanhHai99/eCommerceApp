package dto

import (
	"eCommerce/model"
)

type GetAllOrderDataRes struct {
	Total int16        `json:"total"`
	Page  int16        `json:"page"`
	Data  []model.Order `json:"data"`
}

type GetAllOrderRes struct {
	BaseRes
	Data GetAllOrderDataRes `json:"data"`
}

type GetOneOrderRes struct {
	BaseRes
	Data model.Order `json:"data`
}

type UpdateOrderRes struct {
	BaseRes
	Data model.Order `json:"data`
}

type DeleteOrderRes struct {
	BaseRes
}

type OrderBody struct {
	//Name       string    `json:"name"`
	//Category   uuid.UUID `json:"category"`
	//Detail     string    `json:"detail"`
	//UserManual string    `json:"user_manual"`
	//Price      int16     `json:"price"`
	//CreatedBy  uuid.UUID `json:"created_by"`
}
