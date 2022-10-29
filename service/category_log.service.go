package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllCategoryLog(query map[string]string) (res dto.GetAllCategoryLogRes) {
	dataRes := dto.GetAllCategoryLogDataRes{}
	record, err1 := repository.GetAllCategoryLog()
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

func GetOneCategoryLog(id uuid.UUID) (res dto.GetOneCategoryLogRes) {
	record, err1 := repository.GetOneCategoryLog(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	res.Data = record
	return
}

func UpdateCategoryLog(id uuid.UUID, input map[string]interface{}) (res dto.UpdateCategoryLogRes) {
	item, err1 := repository.UpdateCategoryLog(id, input)
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

func CreateOneCategoryLog(item dto.CategoryLogBody) (res dto.GetOneCategoryLogRes) {
	newCategoryLog := model.CategoryLog{}
	body, _ := json.Marshal(item)
	_ = json.Unmarshal(body, &newCategoryLog)
	err1 := repository.CreateCategoryLog(&newCategoryLog)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	return
}

func DeleteCategoryLog(id uuid.UUID) (res dto.DeleteCategoryLogRes) {
	err := repository.DeleteCategoryLog(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	return
}
