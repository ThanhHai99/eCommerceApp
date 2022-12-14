package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllOrder(query map[string]string) (res dto.GetAllOrderRes) {
	dataRes := dto.GetAllOrderDataRes{}
	record, err1 := repository.GetAllOrder()
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

func GetOneOrder(id uuid.UUID) (res dto.GetOneOrderRes) {
	record, err1 := repository.GetOneOrder(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = &record
	return
}

func UpdateOrder(id uuid.UUID, input map[string]interface{}) (res dto.UpdateOrderRes) {
	item, err1 := repository.UpdateOrder(id, input)
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

func CreateOneOrder(order dto.CreateOrderBody) (res dto.GetOneOrderRes) {
	newOrder := model.Order{}
	body, _ := json.Marshal(order)
	_ = json.Unmarshal(body, &newOrder)
	pre, err1 := repository.CreateOrder(&newOrder)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = pre
	return
}

func DeleteOrder(id uuid.UUID) (res dto.DeleteOrderRes) {
	err := repository.DeleteOrder(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	return
}
