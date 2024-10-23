package routes

import (
	"go-booking/controller"
	"go-booking/middleware"
	"go-booking/service"

	"github.com/gin-gonic/gin"
)

func BookRoute(route *gin.Engine, bookController controller.BookController, jwtService service.JWTService) {
	routes := route.Group("/api/book")
	{
		routes.GET("/", middleware.AuthMiddleware(jwtService), bookController.GetBooks)
		routes.GET("/:id", middleware.AuthMiddleware(jwtService), bookController.GetBookByID)
		routes.POST("/", middleware.AuthMiddleware(jwtService), middleware.RoleMiddleware(jwtService), bookController.CreateBook)
		routes.PATCH("/:id", middleware.AuthMiddleware(jwtService), middleware.RoleMiddleware(jwtService), bookController.UpdateBook)
		routes.DELETE("/:id", middleware.AuthMiddleware(jwtService), middleware.RoleMiddleware(jwtService), bookController.DeleteBook)
	}
}
