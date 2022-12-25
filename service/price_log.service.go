package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllPriceLog(query map[string]string) (res dto.GetAllPriceLogRes) {
	dataRes := dto.GetAllPriceLogDataRes{}
	record, err1 := repository.GetAllPriceLog()
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

func GetOnePriceLog(id uuid.UUID) (res dto.GetOnePriceLogRes) {
	record, err1 := repository.GetOnePriceLog(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = &record
	return
}

func UpdatePriceLog(id uuid.UUID, input map[string]interface{}) (res dto.UpdatePriceLogRes) {
	item, err1 := repository.UpdatePriceLog(id, input)
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

func CreateOnePriceLog(priceLog dto.CreatePriceLogBody) (res dto.GetOnePriceLogRes) {
	newPriceLog := model.PriceLog{}
	body, _ := json.Marshal(priceLog)
	_ = json.Unmarshal(body, &newPriceLog)
	pre, err1 := repository.CreatePriceLog(&newPriceLog)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = pre
	return
}

func DeletePriceLog(id uuid.UUID) (res dto.DeletePriceLogRes) {
	err := repository.DeletePriceLog(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	return
}
