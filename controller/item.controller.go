package controller

import (
	"eCommerce/dto"
	"eCommerce/service"
	"eCommerce/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllItem(c *gin.Context) {
	query := make(map[string]string)
	res := service.GetAllItem(query)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func CreateOneItem(c *gin.Context) {
	input := dto.ItemBody{}
	_ = c.BindJSON(&input)
	res := service.CreateOneItem(input)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

