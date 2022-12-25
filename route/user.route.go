package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"eCommerce/validation"
	"github.com/gin-gonic/gin"
)

func UserRoutes(rootRoute *gin.RouterGroup) {
	userRouter := rootRoute.Group("/user", middleware.ValidateAccessToken())
	{
		userRouter.GET("/", controller.GetAllUser)
		userRouter.GET("/:id", controller.GetOneUser)
		userRouter.PATCH("/:id", controller.UpdateUser)
		userRouter.POST("/", validation.CreateUserBodyValidate(), controller.CreateOneUser)
		userRouter.DELETE("/:id", controller.DeleteUser)
	}
}
