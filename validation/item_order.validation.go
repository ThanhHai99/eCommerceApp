package validation

import (
	"eCommerce/dto"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func CreateItemOrderBodyValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		input := dto.CreateItemOrderBody{}
		_ = c.BindJSON(&input)

		valid := validation.ValidateStruct(&input,
			validation.Field(&input.Amount, validation.Required, validation.Length(2, 10)),
		)

		if valid != nil {
			c.AbortWithStatusJSON(400, valid)
			return
		}
		c.Next()
	}
}
