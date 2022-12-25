package validation

import (
	"eCommerce/dto"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func CreateInvoiceBodyValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		input := dto.CreateInvoiceBody{}
		_ = c.BindJSON(&input)

		valid := validation.ValidateStruct(&input,
			validation.Field(&input.Name, validation.Required, validation.Length(2, 10)),
		)

		if valid != nil {
			c.AbortWithStatusJSON(400, valid)
			return
		}
		c.Next()
	}
}
