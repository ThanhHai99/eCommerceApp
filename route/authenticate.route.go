package route

import (
	"eCommerce/controller"
	"github.com/gin-gonic/gin"
)

func authenticateRoutes(rootRoute *gin.RouterGroup) {
	authenticateRouter := rootRoute.Group("/authentication")
	{
		authenticateRouter.POST("/register", controller.Register)
		authenticateRouter.POST("/login", controller.Login)
		authenticateRouter.GET("/verify/:token", controller.Verify)
	}
}
