package routes

import (
	"go-booking/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine, userController controller.UserController) {
	routes := route.Group("/api/user")
	{
		routes.POST("/register", userController.Register)
	}
}
