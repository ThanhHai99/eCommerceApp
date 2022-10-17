package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
)

func Register(registerBody *dto.RegisterBody) dto.RegisterRes {
	res := dto.RegisterRes{}
	userExisted := repository.CheckExists(registerBody.Username)
	if userExisted == true {
		res.Code = util.FAIL_CODE
		res.Message = "This account is in use"
	} else {
		newUser := model.User{}
		body, _ := json.Marshal(registerBody)
		_ = json.Unmarshal(body, &newUser)
		err1 := repository.CreateNew(&newUser)
		if err1 != nil {
			res.Code = util.FAIL_CODE
			res.Message = "Server error"
		} else {
			res.Code = util.SUCCESS_CODE
			res.Message = "Register successfully"
		}
	}

	return res
}

func Login(loginBody dto.LoginBody) dto.LoginRes {
	res := dto.LoginRes{}

	return res
}

func Verify(token string) dto.VerifyRes {
	res := dto.VerifyRes{}
	record := repository.IsHaveVerifyToken(token)
	if record == nil {
		res.Code = util.FAIL_CODE
		res.Message = "Token is invalid"
	} else {
		err1 := repository.ActiveAccount(record.Username)
		if err1 != nil {
			res.Code = util.FAIL_CODE
			res.Message = "Server error"
		} else {
			res.Code = util.SUCCESS_CODE
			res.Message = "The account is active successfully"
		}
	}

	return res
}
