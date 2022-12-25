package validation

import (
	"eCommerce/dto"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func CreateCategoryBodyValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		input := dto.CreateCategoryBody{}
		_ = c.BindJSON(&input)

		valid := validation.ValidateStruct(&input,
			validation.Field(&input.Name, validation.Required),
		)

		if valid != nil {
			c.AbortWithStatusJSON(400, valid)
			return
		}
		c.Next()
	}
}

func UpdateCategoryBodyValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		input := dto.UpdateCategoryBody{}
		_ = c.BindJSON(&input)

		valid := validation.ValidateStruct(&input,
			validation.Field(&input.Name, validation.NilOrNotEmpty),
		)

		if valid != nil {
			c.AbortWithStatusJSON(400, valid)
			return
		}
		c.Next()
	}
}
