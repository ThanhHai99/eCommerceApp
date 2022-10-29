package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllItemOrder(query map[string]string) (res dto.GetAllItemOrderRes) {
	dataRes := dto.GetAllItemOrderDataRes{}
	record, err1 := repository.GetAllItemOrder()
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

func GetOneItemOrder(id uuid.UUID) (res dto.GetOneItemOrderRes) {
	record, err1 := repository.GetOneItemOrder(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	res.Data = record
	return
}

func UpdateItemOrder(id uuid.UUID, input map[string]interface{}) (res dto.UpdateItemOrderRes) {
	item, err1 := repository.UpdateItemOrder(id, input)
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

func CreateOneItemOrder(item dto.ItemOrderBody) (res dto.GetOneItemOrderRes) {
	newItemOrder := model.ItemOrder{}
	body, _ := json.Marshal(item)
	_ = json.Unmarshal(body, &newItemOrder)
	err1 := repository.CreateItemOrder(&newItemOrder)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	return
}

func DeleteItemOrder(id uuid.UUID) (res dto.DeleteItemOrderRes) {
	err := repository.DeleteItemOrder(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	return
}
