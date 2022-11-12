package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllItem(query map[string]string) (res dto.GetAllItemRes) {
	dataRes := dto.GetAllItemDataRes{}
	record, err1 := repository.GetAllItem()
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

func GetOneItem(id uuid.UUID) (res dto.GetOneItemRes) {
	record, err1 := repository.GetOneItem(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = &record
	return
}

func UpdateItem(id uuid.UUID, input map[string]interface{}) (res dto.UpdateItemRes) {
	item, err1 := repository.UpdateItem(id, input)
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

func CreateOneItem(item dto.ItemBody) (res dto.GetOneItemRes) {
	newItem := model.Item{}
	body, _ := json.Marshal(item)
	_ = json.Unmarshal(body, &newItem)
	pre, err1 := repository.CreateItem(&newItem)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = pre
	return
}

func DeleteItem(id uuid.UUID) (res dto.DeleteItemRes) {
	err := repository.DeleteItem(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	return
}
