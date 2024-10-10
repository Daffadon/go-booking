package routes

import (
	"go-booking/controller"
	"go-booking/service"

	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/api/user")
	{
		routes.POST("/register", userController.Register)
		routes.POST("/login", userController.Login)
		// routes.GET("/profile", middleware.AuthMiddleware(jwtService))
	}
}
