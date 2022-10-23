package service

import (
	"eCommerce/dto"
	"eCommerce/helper"
	"eCommerce/model"
	"eCommerce/repository"
	"eCommerce/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context, registerBody *dto.RegisterBody) dto.RegisterRes {
	res := dto.RegisterRes{}
	userExisted := repository.CheckExists(registerBody.Username)
	if userExisted == true {
		res.Code = util.FAIL_CODE
		res.Message = "This account is in use"
	} else {
		newUser := model.User{}
		passwordHashed, _ := helper.HashPassword(registerBody.Password)
		registerBody.Password = passwordHashed
		newUser.VerifyToken = uuid.New().String()
		body, _ := json.Marshal(registerBody)
		_ = json.Unmarshal(body, &newUser)
		err2 := repository.CreateNew(&newUser)
		if err2 != nil {
			res.Code = util.FAIL_CODE
			res.Message = "Server error"
		} else {
			protocol := "http"
			if c.Request.TLS != nil {
				protocol = "https"
			}
			activeUserUrl := protocol + "://" + c.Request.Host + "/api/verify/" + newUser.VerifyToken
			mailContent := "Click this link to active user: " + activeUserUrl
			err3 := util.SendMail(newUser.Username, "Active account eCommerce", mailContent)
			if err3 == nil {
				res.IsSentMailActive = true
			}

			res.Code = util.SUCCESS_CODE
			res.Message = "Register successfully"
		}
	}

	return res
}

func Login(loginBody dto.LoginBody) dto.LoginRes {
	res := dto.LoginRes{}
	user := repository.GetOneByUsername(loginBody.Username)
	if user == nil {
		res.Code = util.FAIL_CODE
		res.Message = "The username is not exists"
	} else {
		err1 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginBody.Password))
		if err1 == nil {
			accessToken, _ := helper.CreateAccessToken(user.Username)
			res.Code = util.SUCCESS_CODE
			res.Message = "Login successfully"
			res.AccessToken = accessToken
		} else {
			res.Code = util.FAIL_CODE
			res.Message = "The credential is invalid"
		}
	}
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
