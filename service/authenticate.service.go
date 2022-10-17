package service

import (
	"eCommerce/dto"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
)

func Register(registerBody dto.RegisterBody) dto.RegisterRes {
	res := dto.RegisterRes{}
	db, _ := model.Connect()
	defer db.Close()
	err1 := db.Create(registerBody).Error
	if err1 != nil {
		res.Code = util.FAIL_CODE
		res.Message = "Lỗi hệ thống"
	} else {
		res.Code = util.SUCCESS_CODE
		res.Message = "Đăng ký thành công"
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
	if record == nil || record.RowsAffected == 0 {
		res.Code = util.FAIL_CODE
		res.Message = "Token không đúng"
	} else {
		err1 := record.Update(model.User{IsActive: true}).Error
		if err1 != nil {
			res.Code = util.FAIL_CODE
			res.Message = "Lỗi hệ thống"
		} else {
			res.Code = util.SUCCESS_CODE
			res.Message = "Tài khoản đã được kích hoạt"
		}
	}

	return res
}
