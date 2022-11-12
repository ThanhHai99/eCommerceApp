package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllInvoice(query map[string]string) (res dto.GetAllInvoiceRes) {
	dataRes := dto.GetAllInvoiceDataRes{}
	record, err1 := repository.GetAllInvoice()
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

func GetOneInvoice(id uuid.UUID) (res dto.GetOneInvoiceRes) {
	record, err1 := repository.GetOneInvoice(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = &record
	return
}

func UpdateInvoice(id uuid.UUID, input map[string]interface{}) (res dto.UpdateInvoiceRes) {
	item, err1 := repository.UpdateInvoice(id, input)
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

func CreateOneInvoice(item dto.InvoiceBody) (res dto.GetOneInvoiceRes) {
	newInvoice := model.Invoice{}
	body, _ := json.Marshal(item)
	_ = json.Unmarshal(body, &newInvoice)
	pre, err1 := repository.CreateInvoice(&newInvoice)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = pre
	return
}

func DeleteInvoice(id uuid.UUID) (res dto.DeleteInvoiceRes) {
	err := repository.DeleteInvoice(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	return
}
