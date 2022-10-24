package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllItem(query map[string]string) dto.GetAllItemRes {
	res := dto.GetAllItemRes{}
	dataRes := dto.GetAllItemDataRes{}

	record, err1 := repository.GetAllItem()
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	} else {
		dataRes.Data = record
		res.Code = util.SUCCESS_CODE
		res.Message = "Successfully"
		res.Data = dataRes
	}

	return res
}

func GetOneItem(id uuid.UUID) dto.GetOneItemRes {
	res := dto.GetOneItemRes{}
	record, err1 := repository.GetOneItem(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	} else {
		res.Code = util.SUCCESS_CODE
		res.Message = "Successfully"
		res.Data = record
	}

	return res
}

func CreateOneItem(item dto.ItemBody) dto.GetOneItemRes {
	res := dto.GetOneItemRes{}
	newItem := model.Item{}
	body, _ := json.Marshal(item)
	_ = json.Unmarshal(body, &newItem)
	err1 := repository.CreateItem(&newItem)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	} else {
		res.Code = util.SUCCESS_CODE
		res.Message = "Successfully"
	}

	return res
}
