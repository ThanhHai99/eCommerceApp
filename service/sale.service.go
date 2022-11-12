package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllSale(query map[string]string) (res dto.GetAllSaleRes) {
	dataRes := dto.GetAllSaleDataRes{}
	record, err1 := repository.GetAllSale()
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

func GetOneSale(id uuid.UUID) (res dto.GetOneSaleRes) {
	record, err1 := repository.GetOneSale(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = &record
	return
}

func UpdateSale(id uuid.UUID, input map[string]interface{}) (res dto.UpdateSaleRes) {
	item, err1 := repository.UpdateSale(id, input)
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

func CreateOneSale(item dto.SaleBody) (res dto.GetOneSaleRes) {
	newSale := model.Sale{}
	body, _ := json.Marshal(item)
	_ = json.Unmarshal(body, &newSale)
	err1 := repository.CreateSale(&newSale)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	return
}

func DeleteSale(id uuid.UUID) (res dto.DeleteSaleRes) {
	err := repository.DeleteSale(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	return
}
