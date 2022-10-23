package middleware

import (
	"eCommerce/dto"
	"eCommerce/helper"
	"eCommerce/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ValidateAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := dto.BaseRes{}
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.Contains(authHeader, "Bearer") {
			res.Code = util.FAIL_CODE
			res.Message = "Failed to process request"
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		accessToken := strings.Split(strings.TrimSpace(authHeader), " ")
		isAccessTokenValid := helper.ValidateToken(accessToken[1])
		if isAccessTokenValid == true {
			fmt.Println("token dung roi ba")
		} else {
			res.Code = util.FAIL_CODE
			res.Message = "Token is not valid"
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		}
	}
}
