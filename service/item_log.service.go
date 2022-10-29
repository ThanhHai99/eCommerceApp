package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllItemLog(query map[string]string) (res dto.GetAllItemLogRes) {
	dataRes := dto.GetAllItemLogDataRes{}
	record, err1 := repository.GetAllItemLog()
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

func GetOneItemLog(id uuid.UUID) (res dto.GetOneItemLogRes) {
	record, err1 := repository.GetOneItemLog(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	res.Data = record
	return
}

func UpdateItemLog(id uuid.UUID, input map[string]interface{}) (res dto.UpdateItemLogRes) {
	item, err1 := repository.UpdateItemLog(id, input)
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

func CreateOneItemLog(item dto.ItemLogBody) (res dto.GetOneItemLogRes) {
	newItemLog := model.ItemLog{}
	body, _ := json.Marshal(item)
	_ = json.Unmarshal(body, &newItemLog)
	err1 := repository.CreateItemLog(&newItemLog)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	return
}

func DeleteItemLog(id uuid.UUID) (res dto.DeleteItemLogRes) {
	err := repository.DeleteItemLog(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	return
}
