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

func CreateOneSale(sale dto.SaleBody, saleItem dto.SaleItemBody) (res dto.GetOneSaleRes) {
	newSale := model.Sale{}
	newSaleItem := model.SaleItem{}
	body1, _ := json.Marshal(sale)
	body2, _ := json.Marshal(saleItem)
	_ = json.Unmarshal(body1, &newSale)
	_ = json.Unmarshal(body2, &newSaleItem)
	pre, err1 := repository.CreateSale(&newSale)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}

	_, _ = repository.CreateSaleItem(&newSaleItem)

	newSaleLog := model.SaleLog{}
	newSaleLog.Name = pre.Name
	newSaleLog.SaleID = pre.ID
	newSaleLog.SaleItem = saleItem.Item
	newSaleLog.StartDate = pre.StartDate
	newSaleLog.EndDate = pre.EndDate
	newSaleLog.Amount = saleItem.Amount
	newSaleLog.Discount = pre.Discount
	newSaleLog.IsApplied = pre.IsApplied
	newSaleLog.Code = pre.Code
	newSaleLog.CreatedBy = pre.CreatedBy
	_, _ = repository.CreateSaleLog(&newSaleLog)

	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = pre
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
