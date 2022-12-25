package controller

import (
	"eCommerce/dto"
	"eCommerce/service"
	"eCommerce/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func GetAllSaleItem(c *gin.Context) {
	query := make(map[string]string)
	res := service.GetAllSaleItem(query)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetOneSaleItem(c *gin.Context) {
	id := c.Param("id")
	idUid, _ := uuid.Parse(id)
	res := service.GetOneSaleItem(idUid)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func UpdateSaleItem(c *gin.Context) {
	id := c.Param("id")
	idUid, _ := uuid.Parse(id)
	var input map[string]interface{}
	_ = c.BindJSON(&input)
	res := service.UpdateSaleItem(idUid, input)
	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func CreateOneSaleItem(c *gin.Context) {
	input := dto.CreateSaleItemBody{}
	_ = c.BindJSON(&input)
	res := service.CreateOneSaleItem(input)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func DeleteSaleItem(c *gin.Context) {
	id := c.Param("id")
	idUid, _ := uuid.Parse(id)
	res := service.DeleteSaleItem(idUid)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
