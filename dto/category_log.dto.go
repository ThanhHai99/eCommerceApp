package dto

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

type GetAllCategoryLogDataRes struct {
	Total int16               `json:"total"`
	Page  int16               `json:"page"`
	Data  []model.CategoryLog `json:"data"`
}

type GetAllCategoryLogRes struct {
	BaseRes
	Data GetAllCategoryLogDataRes `json:"data"`
}

type GetOneCategoryLogRes struct {
	BaseRes
	Data *model.CategoryLog `json:"data"`
}

type UpdateCategoryLogRes struct {
	BaseRes
	Data *model.CategoryLog `json:"data"`
}

type DeleteCategoryLogRes struct {
	BaseRes
}

type CategoryLogBody struct {
	Name      string    `json:"name"`
	Category  uuid.UUID `json:"category"`
	CreatedBy uuid.UUID `json:"created_by"`
}
