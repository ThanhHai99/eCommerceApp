package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllWarehouse(query map[string]string) (res dto.GetAllWarehouseRes) {
	dataRes := dto.GetAllWarehouseDataRes{}
	record, err1 := repository.GetAllWarehouse()
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

func GetOneWarehouse(id uuid.UUID) (res dto.GetOneWarehouseRes) {
	record, err1 := repository.GetOneWarehouse(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = &record
	return
}

func UpdateWarehouse(id uuid.UUID, input map[string]interface{}) (res dto.UpdateWarehouseRes) {
	item, err1 := repository.UpdateWarehouse(id, input)
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

func CreateOneWarehouse(warehouse dto.WarehouseBody) (res dto.GetOneWarehouseRes) {
	newWarehouse := model.Warehouse{}
	body, _ := json.Marshal(warehouse)
	_ = json.Unmarshal(body, &newWarehouse)
	pre, err1 := repository.CreateWarehouse(&newWarehouse)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}

	newWarehouseLog := model.WarehouseLog{}
	newWarehouseLog.Status = "+"
	//newWarehouseLog.Price =
	newWarehouseLog.WarehouseID = pre.ID
	//newWarehouseLog.Item =
	//newWarehouseLog.Amount =
	//newWarehouseLog.ExpirationDate =
	//newWarehouseLog.ExpirationDate =
	newWarehouseLog.CreatedBy = warehouse.CreatedBy

	_, _ = repository.CreateWarehouseLog(&newWarehouseLog)

	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = pre
	return
}

func DeleteWarehouse(id uuid.UUID) (res dto.DeleteWarehouseRes) {
	err := repository.DeleteWarehouse(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	return
}
