package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

type GetAllCategoryDataRes struct {
	Total int16            `json:"total"`
	Page  int16            `json:"page"`
	Data  []model.Category `json:"data"`
}

type GetAllCategoryRes struct {
	BaseRes
	Data GetAllCategoryDataRes `json:"data"`
}

type GetOneCategoryRes struct {
	BaseRes
	Data *model.Category `json:"data"`
}

type UpdateCategoryRes struct {
	BaseRes
	Data *model.Category `json:"data"`
}

type DeleteCategoryRes struct {
	BaseRes
}

type CategoryBody struct {
	Name      string    `json:"name"`
	CreatedBy uuid.UUID `json:"created_by"`
}
