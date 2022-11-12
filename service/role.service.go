package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/google/uuid"
)

func GetAllRole(query map[string]string) (res dto.GetAllRoleRes) {
	dataRes := dto.GetAllRoleDataRes{}
	record, err1 := repository.GetAllRole()
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

func GetOneRole(id uuid.UUID) (res dto.GetOneRoleRes) {
	record, err1 := repository.GetOneRole(id)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = &record
	return
}

func UpdateRole(id uuid.UUID, input map[string]interface{}) (res dto.UpdateRoleRes) {
	item, err1 := repository.UpdateRole(id, input)
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

func CreateOneRole(role dto.RoleBody) (res dto.GetOneRoleRes) {
	newRole := model.Role{}
	body, _ := json.Marshal(role)
	_ = json.Unmarshal(body, &newRole)
	pre, err1 := repository.CreateRole(&newRole)
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Server error"
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	res.Data = pre
	return
}

func DeleteRole(id uuid.UUID) (res dto.DeleteRoleRes) {
	err := repository.DeleteRole(id)
	if err != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Fail"
		return
	}
	res.Code = util.SUCCESS_CODE
	res.Message = "successfully"
	return
}
