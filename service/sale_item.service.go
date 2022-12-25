package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllSaleItem(query map[string]string) (res dto.GetAllSaleItemRes) {
	dataRes := dto.GetAllSaleItemDataRes{}
	record, err1 := repository.GetAllSaleItem()
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

func GetOneSaleItem(id uuid.UUID) (res dto.GetOneSaleItemRes) {
	record, err1 := repository.GetOneSaleItem(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = &record
	return
}

func UpdateSaleItem(id uuid.UUID, input map[string]interface{}) (res dto.UpdateSaleItemRes) {
	item, err1 := repository.UpdateSaleItem(id, input)
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

func CreateOneSaleItem(saleItem dto.CreateSaleItemBody) (res dto.GetOneSaleItemRes) {
	newSaleItem := model.SaleItem{}
	body, _ := json.Marshal(saleItem)
	_ = json.Unmarshal(body, &newSaleItem)
	pre, err1 := repository.CreateSaleItem(&newSaleItem)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = pre
	return
}

func DeleteSaleItem(id uuid.UUID) (res dto.DeleteSaleItemRes) {
	err := repository.DeleteSaleItem(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	return
}
