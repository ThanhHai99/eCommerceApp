package controller

import (
	"eCommerce/dto"
	"eCommerce/service"
	"eCommerce/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	input := dto.RegisterBody{}
	_ = c.BindJSON(&input)
	res := service.Register(c, &input)
	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func Login(c *gin.Context) {
	input := dto.LoginBody{}
	_ = c.BindJSON(&input)
	res := service.Login(input)
	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func Verify(c *gin.Context) {
	token := c.Param("token")

	res := service.Verify(token)
	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
