package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllWarehouseLog(query map[string]string) (res dto.GetAllWarehouseLogRes) {
	dataRes := dto.GetAllWarehouseLogDataRes{}
	record, err1 := repository.GetAllWarehouseLog()
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

func GetOneWarehouseLog(id uuid.UUID) (res dto.GetOneWarehouseLogRes) {
	record, err1 := repository.GetOneWarehouseLog(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = &record
	return
}

func UpdateWarehouseLog(id uuid.UUID, input map[string]interface{}) (res dto.UpdateWarehouseLogRes) {
	item, err1 := repository.UpdateWarehouseLog(id, input)
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

func CreateOneWarehouseLog(warehouseLog dto.CreateWarehouseLogBody) (res dto.GetOneWarehouseLogRes) {
	newWarehouseLog := model.WarehouseLog{}
	body, _ := json.Marshal(warehouseLog)
	_ = json.Unmarshal(body, &newWarehouseLog)
	pre, err1 := repository.CreateWarehouseLog(&newWarehouseLog)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = pre
	return
}

func DeleteWarehouseLog(id uuid.UUID) (res dto.DeleteWarehouseLogRes) {
	err := repository.DeleteWarehouseLog(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	return
}
