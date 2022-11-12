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
	res.Message = "successfully"
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
	res.Message = "successfully"
	res.Data = &record
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
	res.Message = "successfully"
	res.Data = &item
	return
}

func CreateOneCategory(category dto.CategoryBody) (res dto.GetOneCategoryRes) {
	newCategory := model.Category{}
	body, _ := json.Marshal(category)
	_ = json.Unmarshal(body, &newCategory)
	pre, err1 := repository.CreateCategory(&newCategory)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}

	newCategoryLog := model.CategoryLog{}
	newCategoryLog.Name = newCategory.Name
	newCategoryLog.CreatedBy = newCategory.CreatedBy
	newCategoryLog.CategoryID = pre.ID
	_, _ = repository.CreateCategoryLog(&newCategoryLog)

	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = pre
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
	res.Message = "successfully"
	return
}
