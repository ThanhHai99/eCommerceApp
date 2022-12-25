package controller

import (
	"eCommerce/dto"
	"eCommerce/service"
	"eCommerce/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func GetAllItemLog(c *gin.Context) {
	query := make(map[string]string)
	res := service.GetAllItemLog(query)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetOneItemLog(c *gin.Context) {
	id := c.Param("id")
	idUid, _ := uuid.Parse(id)
	res := service.GetOneItemLog(idUid)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func UpdateItemLog(c *gin.Context) {
	id := c.Param("id")
	idUid, _ := uuid.Parse(id)
	var input map[string]interface{}
	_ = c.BindJSON(&input)
	res := service.UpdateItemLog(idUid, input)
	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func CreateOneItemLog(c *gin.Context) {
	input := dto.CreateItemLogBody{}
	_ = c.BindJSON(&input)
	res := service.CreateOneItemLog(input)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func DeleteItemLog(c *gin.Context) {
	id := c.Param("id")
	idUid, _ := uuid.Parse(id)
	res := service.DeleteItemLog(idUid)

	if res.Code == util.FAIL_CODE {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
