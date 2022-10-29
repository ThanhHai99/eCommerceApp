package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllUser(query map[string]string) (res dto.GetAllUserRes) {
	dataRes := dto.GetAllUserDataRes{}
	record, err1 := repository.GetAllUser()
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

func GetOneUser(id uuid.UUID) (res dto.GetOneUserRes) {
	record, err1 := repository.GetOneUser(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	res.Data = record
	return
}

func UpdateUser(id uuid.UUID, input map[string]interface{}) (res dto.UpdateUserRes) {
	item, err1 := repository.UpdateUser(id, input)
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

func CreateOneUser(item dto.UserBody) (res dto.GetOneUserRes) {
	newUser := model.User{}
	body, _ := json.Marshal(item)
	_ = json.Unmarshal(body, &newUser)
	err1 := repository.CreateUser(&newUser)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	return
}

func DeleteUser(id uuid.UUID) (res dto.DeleteUserRes) {
	err := repository.DeleteUser(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "Successfully"
	return
}
