package controller

import (
	"eCommerce/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	res := service.HealthCheck()
	c.JSON(http.StatusOK, res)
}
