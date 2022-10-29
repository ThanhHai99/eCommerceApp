package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllSaleLog(query map[string]string) (res dto.GetAllSaleLogRes) {
	dataRes := dto.GetAllSaleLogDataRes{}
	record, err1 := repository.GetAllSaleLog()
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

func GetOneSaleLog(id uuid.UUID) (res dto.GetOneSaleLogRes) {
	record, err1 := repository.GetOneSaleLog(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	res.Data = record
	return
}

func UpdateSaleLog(id uuid.UUID, input map[string]interface{}) (res dto.UpdateSaleLogRes) {
	item, err1 := repository.UpdateSaleLog(id, input)
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

func CreateOneSaleLog(item dto.SaleLogBody) (res dto.GetOneSaleLogRes) {
	newSaleLog := model.SaleLog{}
	body, _ := json.Marshal(item)
	_ = json.Unmarshal(body, &newSaleLog)
	err1 := repository.CreateSaleLog(&newSaleLog)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	return
}

func DeleteSaleLog(id uuid.UUID) (res dto.DeleteSaleLogRes) {
	err := repository.DeleteSaleLog(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	return
}
