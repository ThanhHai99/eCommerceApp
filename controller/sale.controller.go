package controller

import (
	"eCommerce/dto"
	"eCommerce/service"
	"eCommerce/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func GetAllSale(c *gin.Context) {
	query := make(map[string]string)
	res := service.GetAllSale(query)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetOneSale(c *gin.Context) {
	id := c.Param("id")
	idUid, _ := uuid.Parse(id)
	res := service.GetOneSale(idUid)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func UpdateSale(c *gin.Context) {
	id := c.Param("id")
	idUid, _ := uuid.Parse(id)
	var input map[string]interface{}
	_ = c.BindJSON(&input)
	res := service.UpdateSale(idUid, input)
	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func CreateOneSale(c *gin.Context) {
	input := dto.SaleBody{}
	_ = c.BindJSON(&input)
	res := service.CreateOneSale(input)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func DeleteSale(c *gin.Context) {
	id := c.Param("id")
	idUid, _ := uuid.Parse(id)
	res := service.DeleteSale(idUid)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
