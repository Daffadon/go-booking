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
		routes.GET("/book", middleware.AuthMiddleware(jwtService), bookController.GetBooks)
		routes.GET("/book/:id", middleware.AuthMiddleware(jwtService), bookController.GetBookByID)
		routes.POST("/book", middleware.AuthMiddleware(jwtService), middleware.RoleMiddleware(jwtService), bookController.CreateBook)
		routes.PATCH("/book/:id", middleware.AuthMiddleware(jwtService), middleware.RoleMiddleware(jwtService), bookController.UpdateBook)
		routes.DELETE("/book/:id", middleware.AuthMiddleware(jwtService), middleware.RoleMiddleware(jwtService), bookController.DeleteBook)
	}
}
