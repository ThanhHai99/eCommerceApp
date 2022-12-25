package dto

import (
	"eCommerce/model"
)

type GetAllCategoryDataRes struct {
	BaseRes
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

type CreateCategoryBody struct {
	Name string `json:"name"`
}

type UpdateCategoryBody struct {
	CreateCategoryBody
}
