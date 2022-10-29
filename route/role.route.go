package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"github.com/gin-gonic/gin"
)

func RoleRoutes(rootRoute *gin.RouterGroup) {
	roleRouter := rootRoute.Group("/role", middleware.ValidateAccessToken())
	{
		roleRouter.GET("/", controller.GetAllCategoryLog)
		roleRouter.GET("/:id", controller.GetOneCategoryLog)
		roleRouter.PATCH("/:id", controller.UpdateCategoryLog)
		roleRouter.POST("/", controller.CreateOneCategoryLog)
		roleRouter.DELETE("/:id", controller.DeleteCategoryLog)
	}
}
