package controller

import (
	"eCommerce/dto"
	"eCommerce/service"
	"eCommerce/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func GetAllWarehouseLog(c *gin.Context) {
	query := make(map[string]string)
	res := service.GetAllWarehouseLog(query)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetOneWarehouseLog(c *gin.Context) {
	id := c.Param("id")
	idUid, _ := uuid.Parse(id)
	res := service.GetOneWarehouseLog(idUid)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func UpdateWarehouseLog(c *gin.Context) {
	id := c.Param("id")
	idUid, _ := uuid.Parse(id)
	var input map[string]interface{}
	_ = c.BindJSON(&input)
	res := service.UpdateWarehouseLog(idUid, input)
	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func CreateOneWarehouseLog(c *gin.Context) {
	input := dto.CreateWarehouseLogBody{}
	_ = c.BindJSON(&input)
	res := service.CreateOneWarehouseLog(input)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func DeleteWarehouseLog(c *gin.Context) {
	id := c.Param("id")
	idUid, _ := uuid.Parse(id)
	res := service.DeleteWarehouseLog(idUid)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
