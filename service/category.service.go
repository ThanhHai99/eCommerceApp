package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllCategory(query map[string]string) (res dto.GetAllCategoryRes) {
	dataRes := dto.GetAllCategoryDataRes{}
	record, err1 := repository.GetAllCategory()
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	dataRes.Data = record
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	res.Data = dataRes
	return
}

func GetOneCategory(id uuid.UUID) (res dto.GetOneCategoryRes) {
	record, err1 := repository.GetOneCategory(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	res.Data = record
	return
}

func UpdateCategory(id uuid.UUID, input map[string]interface{}) (res dto.UpdateCategoryRes) {
	item, err1 := repository.UpdateCategory(id, input)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	res.Data = item
	return
}

func CreateOneCategory(item dto.CategoryBody) (res dto.GetOneCategoryRes) {
	newCategory := model.Category{}
	body, _ := json.Marshal(item)
	_ = json.Unmarshal(body, &newCategory)
	err1 := repository.CreateCategory(&newCategory)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	return
}

func DeleteCategory(id uuid.UUID) (res dto.DeleteCategoryRes) {
	err := repository.DeleteCategory(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	return
}
